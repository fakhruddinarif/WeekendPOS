package service

import (
	"WeekendPOS/app/entity"
	"WeekendPOS/app/gateway/messaging"
	"WeekendPOS/app/model"
	"WeekendPOS/app/model/converter"
	"WeekendPOS/app/repository"
	"context"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"mime/multipart"
)

type UserService struct {
	DB             *gorm.DB
	Log            *logrus.Logger
	Validate       *validator.Validate
	UserRepository *repository.UserRepository
	UserProducer   *messaging.UserProducer
	S3Client       *s3.Client
}

func NewUserService(db *gorm.DB, logger *logrus.Logger, validate *validator.Validate,
	userRepository *repository.UserRepository, userProducer *messaging.UserProducer, s3 *s3.Client) *UserService {
	return &UserService{
		DB:             db,
		Log:            logger,
		Validate:       validate,
		UserRepository: userRepository,
		UserProducer:   userProducer,
		S3Client:       s3,
	}
}

func (s *UserService) Verify(ctx context.Context, request *model.VerifyUserRequest) (*model.Auth, error) {
	tx := s.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	err := s.Validate.Struct(request)
	if err != nil {
		s.Log.Warnf("Invalid request body : %+v", err)
		return nil, fiber.NewError(fiber.StatusBadRequest, "Invalid request body")
	}

	user := new(entity.User)
	if err := s.UserRepository.FindByToken(tx, user, request.Token); err != nil {
		s.Log.Warnf("Failed find user by token : %+v", err)
		return nil, fiber.NewError(fiber.StatusNotFound, "User not found")
	}

	if err := tx.Commit().Error; err != nil {
		s.Log.Warnf("Failed commit transaction : %+v", err)
		return nil, fiber.NewError(fiber.StatusInternalServerError, "Failed commit transaction")
	}

	return &model.Auth{ID: user.ID}, nil
}

func (s *UserService) Create(ctx context.Context, request *model.RegisterUserRequest, fileHeader *multipart.FileHeader) (*model.UserResponse, error) {
	tx := s.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	err := s.Validate.Struct(request)
	if err != nil {
		s.Log.Warnf("Invalid request body : %+v", err)
		return nil, fiber.NewError(fiber.StatusBadRequest, "Invalid request body")
	}

	total, err := s.UserRepository.CountByUsername(tx, request.Username)
	if err != nil {
		s.Log.Warnf("Failed count user from database : %+v", err)
		return nil, fiber.NewError(fiber.StatusInternalServerError, "Failed find user from database")
	}

	if total > 0 {
		s.Log.Warnf("User already exists : %+v", err)
		return nil, fiber.NewError(fiber.StatusConflict, "Username already exists")
	}

	password, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		s.Log.Warnf("Failed to generate bcrype hash : %+v", err)
		return nil, fiber.NewError(fiber.StatusInternalServerError, "Failed to generate bcrype hash")
	}

	var url *string
	if fileHeader != nil {
		url, err = UploadImage("user", s.S3Client, fileHeader)
		if err != nil {
			s.Log.Warnf("Failed to upload image : %+v", err)
			return nil, fiber.NewError(fiber.StatusInternalServerError, "Failed to upload image")
		}
	}

	user := &entity.User{
		Username: request.Username,
		Password: string(password),
		Name:     request.Name,
		Email:    request.Email,
		Phone:    &request.Phone,
		Photo:    url,
	}

	if err := s.UserRepository.Create(tx, user); err != nil {
		s.Log.Warnf("Failed create user to database : %+v", err)
		return nil, fiber.NewError(fiber.StatusInternalServerError, "Failed create user to database")
	}

	if err := tx.Commit().Error; err != nil {
		s.Log.Warnf("Failed commit transaction : %+v", err)
		return nil, fiber.NewError(fiber.StatusInternalServerError, "Failed commit transaction")
	}

	event := converter.UserToEvent(user)
	s.Log.Info("Publishing user created event")
	if err = s.UserProducer.Send(event); err != nil {
		s.Log.Warnf("Failed publish user created event : %+v", err)
		return nil, fiber.NewError(fiber.StatusInternalServerError, "Failed publish user created event")
	}

	return converter.UserToResponse(user), nil
}

func (s *UserService) Login(ctx context.Context, request *model.LoginUserRequest) (*model.UserResponse, error) {
	tx := s.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := s.Validate.Struct(request); err != nil {
		s.Log.Warnf("Invalid request body  : %+v", err)
		return nil, fiber.NewError(fiber.StatusBadRequest, "Invalid request body")
	}

	user := new(entity.User)
	if err := s.UserRepository.FindByUsername(tx, user, request.Username); err != nil {
		s.Log.Warnf("Failed find user by username : %+v", err)
		return nil, fiber.NewError(fiber.StatusUnauthorized, "Username or password is incorrect")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password)); err != nil {
		s.Log.Warnf("Failed to compare user password with bcrype hash : %+v", err)
		return nil, fiber.NewError(fiber.StatusUnauthorized, "Username or password is incorrect")
	}

	token := uuid.New().String()
	user.Token = &token
	if err := s.UserRepository.Update(tx, user); err != nil {
		s.Log.Warnf("Failed save user : %+v", err)
		return nil, fiber.NewError(fiber.StatusInternalServerError, "Failed save token user")
	}

	if err := tx.Commit().Error; err != nil {
		s.Log.Warnf("Failed commit transaction : %+v", err)
		return nil, fiber.NewError(fiber.StatusInternalServerError, "Failed commit transaction")
	}

	event := converter.UserToEvent(user)
	s.Log.Info("Publishing user created event")
	if err := s.UserProducer.Send(event); err != nil {
		s.Log.Warnf("Failed publish user created event : %+v", err)
		return nil, fiber.NewError(fiber.StatusInternalServerError, "Failed publish user created event")
	}

	return converter.UserToTokenResponse(user), nil
}

func (s *UserService) Get(ctx context.Context, request *model.GetUserRequest) (*model.UserResponse, error) {
	tx := s.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := s.Validate.Struct(request); err != nil {
		s.Log.Warnf("Invalid request body : %+v", err)
		return nil, fiber.NewError(fiber.StatusBadRequest, "Invalid request body")
	}

	user := new(entity.User)
	if err := s.UserRepository.FindById(tx, user, request.ID); err != nil {
		s.Log.Warnf("Failed find user by id : %+v", err)
		return nil, fiber.NewError(fiber.StatusNotFound, "User not found")
	}

	if err := tx.Commit().Error; err != nil {
		s.Log.Warnf("Failed commit transaction : %+v", err)
		return nil, fiber.NewError(fiber.StatusInternalServerError, "Failed commit transaction")
	}
	return converter.UserToResponse(user), nil
}

func (s *UserService) Update(ctx context.Context, request *model.UpdateUserRequest, fileHeader *multipart.FileHeader) (*model.UserResponse, error) {
	tx := s.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := s.Validate.Struct(request); err != nil {
		s.Log.Warnf("Invalid request body : %+v", err)
		return nil, fiber.NewError(fiber.StatusBadRequest, "Invalid request body")
	}

	user := new(entity.User)
	if err := s.UserRepository.FindById(tx, user, request.ID); err != nil {
		s.Log.Warnf("Failed find user by id : %+v", err)
		return nil, fiber.NewError(fiber.StatusNotFound, "User not found")
	}

	setIfNotEmpty(&user.Name, request.Name)
	setIfNotEmpty(&user.Email, request.Email)
	setIfNotEmpty(user.Phone, request.Phone)
	setIfNotEmpty(&user.Username, request.Username)
	if request.Password != "" {
		password, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
		if err != nil {
			s.Log.Warnf("Failed to generate bcrype hash : %+v", err)
			return nil, fiber.NewError(fiber.StatusInternalServerError, "Failed to generate bcrype hash")
		}
		user.Password = string(password)
	}
	if fileHeader != nil {
		url, err := UploadImage("user", s.S3Client, fileHeader)
		if err != nil {
			s.Log.Warnf("Failed to upload image : %+v", err)
			return nil, fiber.NewError(fiber.StatusInternalServerError, "Failed to upload image")
		}
		user.Photo = url
	}

	if err := s.UserRepository.Update(tx, user); err != nil {
		s.Log.Warnf("Failed save user : %+v", err)
		return nil, fiber.NewError(fiber.StatusInternalServerError, "Failed update user")
	}

	if err := tx.Commit().Error; err != nil {
		s.Log.Warnf("Failed commit transaction : %+v", err)
		return nil, fiber.NewError(fiber.StatusInternalServerError, "Failed commit transaction")
	}
	event := converter.UserToEvent(user)
	s.Log.Info("Publishing user updated event")
	if err := s.UserProducer.Send(event); err != nil {
		s.Log.Warnf("Failed publish user updated event : %+v", err)
		return nil, fiber.NewError(fiber.StatusInternalServerError, "Failed publish user updated event")
	}
	return converter.UserToResponse(user), nil
}

func (s *UserService) Logout(ctx context.Context, request *model.LogoutUserRequest) (bool, error) {
	tx := s.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := s.Validate.Struct(request); err != nil {
		s.Log.Warnf("Invalid request body : %+v", err)
		return false, fiber.NewError(fiber.StatusBadRequest, "Invalid request body")
	}

	user := new(entity.User)
	if err := s.UserRepository.FindById(tx, user, request.ID); err != nil {
		s.Log.Warnf("Failed find user by id : %+v", err)
		return false, fiber.NewError(fiber.StatusNotFound, "User not found")
	}

	user.Token = nil
	if err := s.UserRepository.Update(tx, user); err != nil {
		s.Log.Warnf("Failed save user : %+v", err)
		return false, fiber.NewError(fiber.StatusInternalServerError, "Failed update user")
	}

	if err := tx.Commit().Error; err != nil {
		s.Log.Warnf("Failed commit transaction : %+v", err)
		return false, fiber.NewError(fiber.StatusInternalServerError, "Failed commit transaction")
	}

	event := converter.UserToEvent(user)
	s.Log.Info("Publishing user updated event")
	if err := s.UserProducer.Send(event); err != nil {
		s.Log.Warnf("Failed publish user updated event : %+v", err)
		return false, fiber.NewError(fiber.StatusInternalServerError, "Failed publish user updated event")
	}
	return true, nil
}

func setIfNotEmpty(target *string, value string) {
	if value != "" {
		*target = value
	}
}

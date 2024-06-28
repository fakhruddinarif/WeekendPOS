package service

import (
	"WeekendPOS/app/entity"
	"WeekendPOS/app/gateway/messaging"
	"WeekendPOS/app/model"
	"WeekendPOS/app/model/converter"
	"WeekendPOS/app/repository"
	"context"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type EmployeeService struct {
	DB                 *gorm.DB
	Log                *logrus.Logger
	Validate           *validator.Validate
	EmployeeRepository *repository.EmployeeRepository
	EmployeeProducer   *messaging.EmployeeProducer
}

func NewEmployeeService(db *gorm.DB, log *logrus.Logger, validate *validator.Validate, employeeRepository *repository.EmployeeRepository, employeeProducer *messaging.EmployeeProducer) *EmployeeService {
	return &EmployeeService{
		DB:                 db,
		Log:                log,
		Validate:           validate,
		EmployeeRepository: employeeRepository,
		EmployeeProducer:   employeeProducer,
	}
}

func (s *EmployeeService) Create(ctx context.Context, request *model.CreateEmployeeRequest) (*model.EmployeeResponse, error) {
	tx := s.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := s.Validate.Struct(request); err != nil {
		s.Log.WithError(err).Error("validation error request body.")
		return nil, fiber.ErrBadRequest
	}

	password, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		s.Log.WithError(err).Error("failed to hash password.")
		return nil, fiber.ErrInternalServerError
	}

	employee := &entity.Employee{
		Name:     request.Name,
		Email:    request.Email,
		Username: request.Username,
		Password: string(password),
		Phone:    request.Phone,
		Address:  request.Address,
		Photo:    request.Photo,
		UserId:   request.UserID,
	}

	if err := s.EmployeeRepository.Create(tx, employee); err != nil {
		s.Log.WithError(err).Error("failed to create employee.")
		return nil, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		s.Log.WithError(err).Error("failed to commit transaction.")
		return nil, fiber.ErrInternalServerError
	}

	event := converter.EmployeeToEvent(employee)
	if err := s.EmployeeProducer.Send(event); err != nil {
		s.Log.WithError(err).Error("failed to send message.")
		return nil, fiber.ErrInternalServerError
	}
	return converter.EmployeeToResponse(employee), nil
}

func (s *EmployeeService) Update(ctx context.Context, request *model.UpdateEmployeeRequest) (*model.EmployeeResponse, error) {
	tx := s.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := s.Validate.Struct(request); err != nil {
		s.Log.WithError(err).Error("validation error request body.")
		return nil, fiber.ErrBadRequest
	}

	employee := new(entity.Employee)
	if err := s.EmployeeRepository.FindById(tx, employee, request.ID); err != nil {
		s.Log.WithError(err).Error("failed to find employee.")
		return nil, fiber.ErrNotFound
	}

	employee.Name = request.Name
	employee.Email = request.Email
	employee.Phone = request.Phone
	employee.Address = request.Address
	employee.Photo = request.Photo

	if err := s.EmployeeRepository.Update(tx, employee); err != nil {
		s.Log.WithError(err).Error("failed to update employee.")
		return nil, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		s.Log.WithError(err).Error("failed to commit transaction.")
		return nil, fiber.ErrInternalServerError
	}

	event := converter.EmployeeToEvent(employee)
	if err := s.EmployeeProducer.Send(event); err != nil {
		s.Log.WithError(err).Error("failed to send message.")
		return nil, fiber.ErrInternalServerError
	}
	return converter.EmployeeToResponse(employee), nil
}

func (s *EmployeeService) List(ctx context.Context, request *model.ListEmployeeRequest) ([]model.EmployeeResponse, int64, error) {
	tx := s.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := s.Validate.Struct(request); err != nil {
		s.Log.WithError(err).Error("validation error request body.")
		return nil, 0, fiber.ErrBadRequest
	}

	employees, total, err := s.EmployeeRepository.Search(tx, request)
	if err != nil {
		s.Log.WithError(err).Error("failed to list employee.")
		return nil, 0, fiber.ErrInternalServerError
	}

	responses := make([]model.EmployeeResponse, len(employees))
	for i, employee := range employees {
		responses[i] = *converter.EmployeeToResponse(&employee)
	}

	if err := tx.Commit().Error; err != nil {
		s.Log.WithError(err).Error("failed to commit transaction.")
		return nil, 0, fiber.ErrInternalServerError
	}

	return responses, total, nil
}

func (s *EmployeeService) Get(ctx context.Context, request *model.GetEmployeeRequest) (*model.EmployeeResponse, error) {
	tx := s.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := s.Validate.Struct(request); err != nil {
		s.Log.WithError(err).Error("validation error request body.")
		return nil, fiber.ErrBadRequest
	}

	employee := new(entity.Employee)
	if err := s.EmployeeRepository.FindById(tx, employee, request.ID); err != nil {
		s.Log.WithError(err).Error("failed to find employee.")
		return nil, fiber.ErrNotFound
	}

	if err := tx.Commit().Error; err != nil {
		s.Log.WithError(err).Error("failed to commit transaction.")
		return nil, fiber.ErrInternalServerError
	}

	return converter.EmployeeToResponse(employee), nil
}

func (s *EmployeeService) Delete(ctx context.Context, request *model.DeleteEmployeeRequest) error {
	tx := s.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := s.Validate.Struct(request); err != nil {
		s.Log.WithError(err).Error("validation error request body.")
		return fiber.ErrBadRequest
	}

	employee := new(entity.Employee)
	if err := s.EmployeeRepository.FindById(tx, employee, request.ID); err != nil {
		s.Log.WithError(err).Error("failed to find employee.")
		return fiber.ErrNotFound
	}

	if err := s.EmployeeRepository.Delete(tx, employee); err != nil {
		s.Log.WithError(err).Error("failed to delete employee.")
		return fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		s.Log.WithError(err).Error("failed to commit transaction.")
		return fiber.ErrInternalServerError
	}

	event := converter.EmployeeToEvent(employee)
	if err := s.EmployeeProducer.Send(event); err != nil {
		s.Log.WithError(err).Error("failed to send message.")
		return fiber.ErrInternalServerError
	}
	return nil
}

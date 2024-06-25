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
	"gorm.io/gorm"
)

type CategoryService struct {
	DB                 *gorm.DB
	Log                *logrus.Logger
	Validate           *validator.Validate
	CategoryRepository *repository.CategoryRepository
	CategoryProducer   *messaging.CategoryProducer
}

func NewCategoryService(db *gorm.DB, log *logrus.Logger, validate *validator.Validate, categoryRepository *repository.CategoryRepository, categoryProducer *messaging.CategoryProducer) *CategoryService {
	return &CategoryService{
		DB:                 db,
		Log:                log,
		Validate:           validate,
		CategoryRepository: categoryRepository,
		CategoryProducer:   categoryProducer,
	}
}

func (s *CategoryService) Create(ctx context.Context, request *model.CreateCategoryRequest) (*model.CategoryResponse, error) {
	tx := s.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := s.Validate.Struct(request); err != nil {
		s.Log.WithError(err).Error("validation error request body.")
		return nil, fiber.ErrBadRequest
	}

	category := &entity.Category{
		Name:   request.Name,
		UserId: request.UserID,
	}

	if err := s.CategoryRepository.Create(tx, category); err != nil {
		s.Log.WithError(err).Error("failed to create category.")
		return nil, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		s.Log.WithError(err).Error("failed to commit transaction.")
		return nil, fiber.ErrInternalServerError
	}

	event := converter.CategoryToEvent(category)
	if err := s.CategoryProducer.Send(event); err != nil {
		s.Log.WithError(err).Error("failed to send message.")
		return nil, fiber.ErrInternalServerError
	}

	return converter.CategoryToResponse(category), nil
}

func (s *CategoryService) Search(ctx context.Context, request *model.SearchCategoryRequest) ([]model.CategoryResponse, int64, error) {
	tx := s.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := s.Validate.Struct(request); err != nil {
		s.Log.WithError(err).Error("validation error request body.")
		return nil, 0, fiber.ErrBadRequest
	}

	categories, total, err := s.CategoryRepository.Search(tx, request)
	if err != nil {
		s.Log.WithError(err).Error("failed to search category.")
		return nil, 0, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		s.Log.WithError(err).Error("failed to commit transaction.")
		return nil, 0, fiber.ErrInternalServerError
	}

	responses := make([]model.CategoryResponse, len(categories))
	for i, category := range categories {
		responses[i] = *converter.CategoryToResponse(category)
	}

	return responses, total, nil

}

func (s *CategoryService) Get(ctx context.Context, request *model.GetCategoryRequest) (*model.CategoryResponse, error) {
	tx := s.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := s.Validate.Struct(request); err != nil {
		s.Log.WithError(err).Error("validation error request body.")
		return nil, fiber.ErrBadRequest
	}

	category := new(entity.Category)
	if err := s.CategoryRepository.FindByIdAndUserId(tx, category, request.ID, request.UserID); err != nil {
		s.Log.WithError(err).Error("failed to getting category.")
		return nil, fiber.ErrNotFound
	}

	if err := tx.Commit().Error; err != nil {
		s.Log.WithError(err).Error("failed to commit transaction.")
		return nil, fiber.ErrInternalServerError
	}

	return converter.CategoryToResponse(category), nil
}

func (s *CategoryService) Update(ctx context.Context, request *model.UpdateCategoryRequest) (*model.CategoryResponse, error) {
	tx := s.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := s.Validate.Struct(request); err != nil {
		s.Log.WithError(err).Error("validation error request body.")
		return nil, fiber.ErrBadRequest
	}

	category := new(entity.Category)
	if err := s.CategoryRepository.FindByIdAndUserId(tx, category, request.ID, request.UserID); err != nil {
		s.Log.WithError(err).Error("failed to getting category.")
		return nil, fiber.ErrNotFound
	}

	category.Name = request.Name

	if err := s.CategoryRepository.Update(tx, category); err != nil {
		s.Log.WithError(err).Error("failed to update category.")
		return nil, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		s.Log.WithError(err).Error("failed to commit transaction.")
		return nil, fiber.ErrInternalServerError
	}

	event := converter.CategoryToEvent(category)
	if err := s.CategoryProducer.Send(event); err != nil {
		s.Log.WithError(err).Error("failed to publishing category.")
		return nil, fiber.ErrInternalServerError
	}

	return converter.CategoryToResponse(category), nil
}

func (s *CategoryService) Delete(ctx context.Context, request *model.DeleteCategoryRequest) error {
	tx := s.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := s.Validate.Struct(request); err != nil {
		s.Log.WithError(err).Error("validation error request body.")
		return fiber.ErrBadRequest
	}

	category := new(entity.Category)
	if err := s.CategoryRepository.FindByIdAndUserId(tx, category, request.ID, request.UserID); err != nil {
		s.Log.WithError(err).Error("failed to getting category.")
		return fiber.ErrNotFound
	}

	if err := s.CategoryRepository.Delete(tx, category); err != nil {
		s.Log.WithError(err).Error("failed to delete category.")
		return fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		s.Log.WithError(err).Error("failed to commit transaction.")
		return fiber.ErrInternalServerError
	}

	return nil
}

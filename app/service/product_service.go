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

type ProductService struct {
	DB                 *gorm.DB
	Log                *logrus.Logger
	Validate           *validator.Validate
	ProductRepository  *repository.ProductRepository
	CategoryRepository *repository.CategoryRepository
	ProductProducer    *messaging.ProductProducer
}

func NewProductService(db *gorm.DB, log *logrus.Logger, validate *validator.Validate, productRepository *repository.ProductRepository, categoryRepository *repository.CategoryRepository, productProducer *messaging.ProductProducer) *ProductService {
	return &ProductService{
		DB:                 db,
		Log:                log,
		Validate:           validate,
		ProductRepository:  productRepository,
		CategoryRepository: categoryRepository,
		ProductProducer:    productProducer,
	}
}

func (s *ProductService) Create(ctx context.Context, request *model.CreateProductRequest) (*model.ProductResponse, error) {
	tx := s.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := s.Validate.Struct(request); err != nil {
		s.Log.WithError(err).Error("validation error request body.")
		return nil, fiber.ErrBadRequest
	}

	product := &entity.Product{
		SKU:        request.SKU,
		Name:       request.Name,
		CategoryId: request.CategoryID,
		BuyPrice:   request.BuyPrice,
		SellPrice:  request.SellPrice,
		Stock:      request.Stock,
		Photo:      request.Photo,
		UserId:     request.UserID,
	}

	if err := s.ProductRepository.Create(tx, product); err != nil {
		s.Log.WithError(err).Error("failed to create product.")
		return nil, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		s.Log.WithError(err).Error("failed to commit transaction.")
		return nil, fiber.ErrInternalServerError
	}

	event := converter.ProductToEvent(product)
	if err := s.ProductProducer.Send(event); err != nil {
		s.Log.WithError(err).Error("failed to send message.")
		return nil, fiber.ErrInternalServerError
	}
	return converter.ProductToResponse(product), nil
}

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

	category := new(entity.Category)
	if err := s.CategoryRepository.FindById(tx, category, request.CategoryID); err != nil {
		s.Log.WithError(err).Error("failed to find category.")
		return nil, fiber.ErrNotFound
	}
	product.Category = *category

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

func (s *ProductService) Get(ctx context.Context, request *model.GetProductRequest) (*model.ProductResponse, error) {
	tx := s.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := s.Validate.Struct(request); err != nil {
		s.Log.WithError(err).Error("validation error request body.")
		return nil, fiber.ErrBadRequest
	}
	product := new(entity.Product)
	if err := s.ProductRepository.FindById(tx, product, request.ID, request.UserID); err != nil {
		s.Log.WithError(err).Error("failed to find product.")
		return nil, fiber.ErrNotFound
	}
	category := new(entity.Category)
	if err := s.CategoryRepository.FindById(tx, category, product.CategoryId); err != nil {
		s.Log.WithError(err).Error("failed to find category.")
		return nil, fiber.ErrNotFound
	}
	product.Category = *category

	if err := tx.Commit().Error; err != nil {
		s.Log.WithError(err).Error("failed to commit transaction.")
		return nil, fiber.ErrInternalServerError
	}
	return converter.ProductToResponse(product), nil
}

func (s *ProductService) List(ctx context.Context, request *model.SearchProductRequest) ([]model.ProductResponse, int64, error) {
	tx := s.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := s.Validate.Struct(request); err != nil {
		s.Log.WithError(err).Error("validation error request body.")
		return nil, 0, fiber.ErrBadRequest
	}

	products, total, err := s.ProductRepository.Search(tx, request)

	if err != nil {
		s.Log.WithError(err).Error("failed to search product.")
		return nil, 0, fiber.ErrInternalServerError
	}

	responses := make([]model.ProductResponse, len(products))
	for i, product := range products {
		category := new(entity.Category)
		if err := s.CategoryRepository.FindById(tx, category, product.CategoryId); err != nil {
			s.Log.WithError(err).Error("failed to find category.")
			return nil, 0, fiber.ErrNotFound
		}
		product.Category = *category
		responses[i] = *converter.ProductToResponse(&product)
	}

	if err := tx.Commit().Error; err != nil {
		s.Log.WithError(err).Error("failed to commit transaction.")
		return nil, 0, fiber.ErrInternalServerError
	}
	return responses, total, nil
}

func (s *ProductService) Update(ctx context.Context, request *model.UpdateProductRequest) (*model.ProductResponse, error) {
	tx := s.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := s.Validate.Struct(request); err != nil {
		s.Log.WithError(err).Error("validation error request body.")
		return nil, fiber.ErrBadRequest
	}

	product := new(entity.Product)
	if err := s.ProductRepository.FindById(tx, product, request.ID, request.UserID); err != nil {
		s.Log.WithError(err).Error("failed to find product.")
		return nil, fiber.ErrNotFound
	}

	product.SKU = request.SKU
	product.Name = request.Name
	product.CategoryId = request.CategoryID
	product.BuyPrice = request.BuyPrice
	product.SellPrice = request.SellPrice
	product.Stock = request.Stock
	product.Photo = request.Photo

	category := new(entity.Category)
	if err := s.CategoryRepository.FindById(tx, category, product.CategoryId); err != nil {
		s.Log.WithError(err).Error("failed to find category.")
		return nil, fiber.ErrNotFound
	}
	product.Category = *category

	if err := s.ProductRepository.Update(tx, product); err != nil {
		s.Log.WithError(err).Error("failed to update product.")
		return nil, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		s.Log.WithError(err).Error("failed to commit transaction.")
		return nil, fiber.ErrInternalServerError
	}
	return converter.ProductToResponse(product), nil
}

func (s *ProductService) Delete(ctx context.Context, request *model.DeleteProductRequest) error {
	tx := s.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := s.Validate.Struct(request); err != nil {
		s.Log.WithError(err).Error("validation error request body.")
		return fiber.ErrBadRequest
	}

	product := new(entity.Product)
	if err := s.ProductRepository.FindById(tx, product, request.ID, request.UserID); err != nil {
		s.Log.WithError(err).Error("failed to find product.")
		return fiber.ErrNotFound
	}

	if err := s.ProductRepository.Delete(tx, product); err != nil {
		s.Log.WithError(err).Error("failed to delete product.")
		return fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		s.Log.WithError(err).Error("failed to commit transaction.")
		return fiber.ErrInternalServerError
	}

	return nil
}

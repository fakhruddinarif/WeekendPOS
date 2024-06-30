package service

import (
	"WeekendPOS/app/gateway/messaging"
	"WeekendPOS/app/model"
	"WeekendPOS/app/repository"
	"context"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type TransactionService struct {
	DB                    *gorm.DB
	Log                   *logrus.Logger
	Validate              *validator.Validate
	TransactionRepository *repository.TransactionRepository
	EmployeeRepository    *repository.EmployeeRepository
	ProductRepository     *repository.ProductRepository
	TransactionProducer   *messaging.TransactionProducer
}

func NewTransactionService(db *gorm.DB, log *logrus.Logger, validate *validator.Validate, transactionRepository *repository.TransactionRepository, employeeRepository *repository.EmployeeRepository, productRepository *repository.ProductRepository, transactionProducer *messaging.TransactionProducer) *TransactionService {
	return &TransactionService{
		DB:                    db,
		Log:                   log,
		Validate:              validate,
		TransactionRepository: transactionRepository,
		EmployeeRepository:    employeeRepository,
		ProductRepository:     productRepository,
		TransactionProducer:   transactionProducer,
	}
}

func (s *TransactionService) Create(ctx context.Context, request *model.CreateTransactionRequest) (*model.TransactionResponse, error) {
	tx := s.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := s.Validate.Struct(request); err != nil {
		s.Log.WithError(err).Error("validation error request body.")
		return nil, fiber.ErrBadRequest
	}
}

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
	"time"
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

	transaction := &entity.Transaction{
		Customer:   request.Customer,
		Date:       time.Now().Format("2006-01-02 15:04:05"),
		EmployeeId: request.EmployeeID,
		UserId:     request.UserID,
	}

	if err := s.TransactionRepository.Create(tx, transaction); err != nil {
		s.Log.WithError(err).Error("failed to create transaction.")
		return nil, fiber.ErrInternalServerError
	}

	if err := s.TransactionRepository.LastInsertedId(tx, transaction); err != nil {
		s.Log.WithError(err).Error("failed to get last inserted id.")
		return nil, fiber.ErrInternalServerError
	}

	total := 0.00
	transactionDetails := make([]entity.DetailTransaction, 0)

	for _, detail := range request.Products {
		product := new(entity.Product)
		if err := s.ProductRepository.FindById(tx, product, detail.ProductID, request.UserID); err != nil {
			s.Log.WithError(err).Error("failed to find product.")
			return nil, fiber.ErrBadRequest
		}

		total += product.SellPrice * float64(detail.Amount)

		if product.Stock < detail.Amount {
			s.Log.Error("stock not enough.")
			return nil, fiber.ErrBadRequest
		}

		transactionDetail := &entity.DetailTransaction{
			TransactionId: transaction.ID,
			ProductId:     detail.ProductID,
			Amount:        detail.Amount,
			Price:         product.SellPrice,
		}

		if err := s.TransactionRepository.CreateDetailTransaction(tx, transactionDetail); err != nil {
			s.Log.WithError(err).Error("failed to create detail transaction.")
			return nil, fiber.ErrInternalServerError
		}

		transactionDetails = append(transactionDetails, *transactionDetail)

		product.Stock -= detail.Amount
		if err := s.ProductRepository.Update(tx, product); err != nil {
			s.Log.WithError(err).Error("failed to update product.")
			return nil, fiber.ErrInternalServerError
		}
	}

	transaction.Total = total

	if err := tx.Commit().Error; err != nil {
		s.Log.WithError(err).Error("failed to commit transaction.")
		return nil, fiber.ErrInternalServerError
	}

	event := converter.TransactionToEvent(transaction)
	if err := s.TransactionProducer.Send(event); err != nil {
		s.Log.WithError(err).Error("failed to send message.")
		return nil, fiber.ErrInternalServerError
	}

	return converter.TransactionToResponse(transaction, transactionDetails), nil
}

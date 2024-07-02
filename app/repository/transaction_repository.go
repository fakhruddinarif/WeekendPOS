package repository

import (
	"WeekendPOS/app/entity"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type TransactionRepository struct {
	Repository[entity.Transaction]
	Log *logrus.Logger
}

func NewTransactionRepository(log *logrus.Logger) *TransactionRepository {
	return &TransactionRepository{
		Log: log,
	}
}

func (r *TransactionRepository) CreateDetailTransaction(db *gorm.DB, detail *entity.DetailTransaction) error {
	return db.Create(detail).Error
}

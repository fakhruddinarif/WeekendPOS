package repository

import (
	"WeekendPOS/app/entity"
	"github.com/sirupsen/logrus"
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

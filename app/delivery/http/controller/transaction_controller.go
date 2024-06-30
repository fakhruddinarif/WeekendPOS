package controller

import (
	"WeekendPOS/app/service"
	"github.com/sirupsen/logrus"
)

type TransactionController struct {
	Service *service.TransactionService
	Log     *logrus.Logger
}

func NewTransactionController(service *service.TransactionService, log *logrus.Logger) *TransactionController {
	return &TransactionController{
		Service: service,
		Log:     log,
	}
}

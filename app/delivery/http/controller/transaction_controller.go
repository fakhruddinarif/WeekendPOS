package controller

import (
	"WeekendPOS/app/delivery/http/middleware"
	"WeekendPOS/app/model"
	"WeekendPOS/app/service"
	"github.com/gofiber/fiber/v2"
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

func (c *TransactionController) Create(ctx *fiber.Ctx) error {
	auth := middleware.GetUser(ctx)

	request := new(model.CreateTransactionRequest)
	if err := ctx.BodyParser(request); err != nil {
		c.Log.WithError(err).Error("Failed to parse request body")
		return fiber.ErrBadRequest
	}

	request.UserID = auth.ID

	response, err := c.Service.Create(ctx.UserContext(), request)
	if err != nil {
		c.Log.WithError(err).Error("Failed to create transaction")
		return err
	}

	return ctx.JSON(model.WebResponse[*model.TransactionResponse]{Data: response})
}

package controller

import (
	"WeekendPOS/app/delivery/http/middleware"
	"WeekendPOS/app/model"
	"WeekendPOS/app/service"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type ProductController struct {
	Service *service.ProductService
	Log     *logrus.Logger
}

func NewProductController(service *service.ProductService, log *logrus.Logger) *ProductController {
	return &ProductController{
		Service: service,
		Log:     log,
	}
}

func (c *ProductController) Create(ctx *fiber.Ctx) error {
	auth := middleware.GetUser(ctx)

	request := new(model.CreateProductRequest)
	if err := ctx.BodyParser(request); err != nil {
		c.Log.WithError(err).Error("Failed to parse request body")
		return fiber.ErrBadRequest
	}
	request.UserID = auth.ID

	response, err := c.Service.Create(ctx.UserContext(), request)
	if err != nil {
		c.Log.WithError(err).Error("Failed to create product")
		return err
	}
	return ctx.JSON(model.WebResponse[*model.ProductResponse]{Data: response})
}

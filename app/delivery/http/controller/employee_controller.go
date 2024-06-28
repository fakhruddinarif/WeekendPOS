package controller

import (
	"WeekendPOS/app/delivery/http/middleware"
	"WeekendPOS/app/model"
	"WeekendPOS/app/service"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"math"
)

type EmployeeController struct {
	Service *service.EmployeeService
	Log     *logrus.Logger
}

func NewEmployeeController(service *service.EmployeeService, log *logrus.Logger) *EmployeeController {
	return &EmployeeController{
		Service: service,
		Log:     log,
	}
}

func (c *EmployeeController) Create(ctx *fiber.Ctx) error {
	auth := middleware.GetUser(ctx)

	request := new(model.CreateEmployeeRequest)
	if err := ctx.BodyParser(request); err != nil {
		c.Log.WithError(err).Error("Failed to parse request body")
		return fiber.ErrBadRequest
	}
	request.UserID = auth.ID

	response, err := c.Service.Create(ctx.UserContext(), request)
	if err != nil {
		c.Log.WithError(err).Error("Failed to create employee")
		return err
	}
	return ctx.JSON(model.WebResponse[*model.EmployeeResponse]{Data: response})
}

func (c *EmployeeController) Update(ctx *fiber.Ctx) error {
	auth := middleware.GetUser(ctx)

	request := new(model.UpdateEmployeeRequest)
	if err := ctx.BodyParser(request); err != nil {
		c.Log.WithError(err).Error("Failed to parse request body")
		return fiber.ErrBadRequest
	}
	request.UserID = auth.ID
	request.ID = ctx.Params("id")

	response, err := c.Service.Update(ctx.UserContext(), request)
	if err != nil {
		c.Log.WithError(err).Error("Failed to update employee")
		return err
	}
	return ctx.JSON(model.WebResponse[*model.EmployeeResponse]{Data: response})
}

func (c *EmployeeController) List(ctx *fiber.Ctx) error {
	auth := middleware.GetUser(ctx)

	request := &model.ListEmployeeRequest{
		UserID:   auth.ID,
		Name:     ctx.Query("name"),
		Email:    ctx.Query("email"),
		Username: ctx.Query("username"),
		Page:     ctx.QueryInt("page", 1),
		Size:     ctx.QueryInt("size", 10),
	}

	response, total, err := c.Service.List(ctx.UserContext(), request)
	if err != nil {
		c.Log.WithError(err).Error("Failed to list employee")
		return err
	}

	paging := &model.PageMetadata{
		Page:      request.Page,
		Size:      request.Size,
		TotalItem: total,
		TotalPage: int64(math.Ceil(float64(total) / float64(request.Size))),
	}

	return ctx.JSON(model.WebResponse[[]model.EmployeeResponse]{Data: response, Paging: paging})
}

func (c *EmployeeController) Get(ctx *fiber.Ctx) error {
	auth := middleware.GetUser(ctx)

	request := &model.GetEmployeeRequest{
		ID:     ctx.Params("id"),
		UserID: auth.ID,
	}

	response, err := c.Service.Get(ctx.UserContext(), request)
	if err != nil {
		c.Log.WithError(err).Error("Failed to get employee")
		return err
	}
	return ctx.JSON(model.WebResponse[*model.EmployeeResponse]{Data: response})
}

func (c *EmployeeController) Delete(ctx *fiber.Ctx) error {
	auth := middleware.GetUser(ctx)

	request := &model.DeleteEmployeeRequest{
		ID:     ctx.Params("id"),
		UserID: auth.ID,
	}

	err := c.Service.Delete(ctx.UserContext(), request)
	if err != nil {
		c.Log.WithError(err).Error("Failed to delete employee")
		return err
	}
	return ctx.JSON(model.WebResponse[bool]{Data: true})
}

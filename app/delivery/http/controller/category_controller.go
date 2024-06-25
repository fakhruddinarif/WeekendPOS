package controller

import (
	"WeekendPOS/app/delivery/http/middleware"
	"WeekendPOS/app/model"
	"WeekendPOS/app/service"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"math"
)

type CategoryController struct {
	Service *service.CategoryService
	Log     *logrus.Logger
}

func NewCategoryController(service *service.CategoryService, log *logrus.Logger) *CategoryController {
	return &CategoryController{
		Service: service,
		Log:     log,
	}
}

func (c *CategoryController) Create(ctx *fiber.Ctx) error {
	auth := middleware.GetUser(ctx)

	request := new(model.CreateCategoryRequest)
	if err := ctx.BodyParser(request); err != nil {
		c.Log.WithError(err).Error("Failed to parse request body")
		return fiber.ErrBadRequest
	}
	request.UserID = auth.ID

	response, err := c.Service.Create(ctx.UserContext(), request)
	if err != nil {
		c.Log.WithError(err).Error("Failed to create category")
		return err
	}
	return ctx.JSON(model.WebResponse[*model.CategoryResponse]{Data: response})
}

func (c *CategoryController) Get(ctx *fiber.Ctx) error {
	auth := middleware.GetUser(ctx)

	request := &model.GetCategoryRequest{
		ID:     ctx.Params("id"),
		UserID: auth.ID,
	}
	response, err := c.Service.Get(ctx.UserContext(), request)
	if err != nil {
		c.Log.WithError(err).Error("Failed to get category")
		return err
	}
	return ctx.JSON(model.WebResponse[*model.CategoryResponse]{Data: response})
}

func (c *CategoryController) List(ctx *fiber.Ctx) error {
	auth := middleware.GetUser(ctx)

	request := &model.SearchCategoryRequest{
		UserID: auth.ID,
		Name:   ctx.Query("name", ""),
		Page:   ctx.QueryInt("page", 1),
		Size:   ctx.QueryInt("size", 10),
	}
	response, total, err := c.Service.Search(ctx.UserContext(), request)
	if err != nil {
		c.Log.WithError(err).Error("Failed to search category")
		return err
	}
	paging := &model.PageMetadata{
		Page:      request.Page,
		Size:      request.Size,
		TotalItem: total,
		TotalPage: int64(math.Ceil(float64(total) / float64(request.Size))),
	}
	return ctx.JSON(model.WebResponse[[]model.CategoryResponse]{Data: response, Paging: paging})
}

func (c *CategoryController) Update(ctx *fiber.Ctx) error {
	auth := middleware.GetUser(ctx)

	request := new(model.UpdateCategoryRequest)
	if err := ctx.BodyParser(request); err != nil {
		c.Log.WithError(err).Error("Failed to parse request body")
		return fiber.ErrBadRequest
	}
	request.UserID = auth.ID
	request.ID = ctx.Params("id")
	response, err := c.Service.Update(ctx.UserContext(), request)
	if err != nil {
		c.Log.WithError(err).Error("Failed to update category")
		return err

	}
	return ctx.JSON(model.WebResponse[*model.CategoryResponse]{Data: response})
}

func (c *CategoryController) Delete(ctx *fiber.Ctx) error {
	auth := middleware.GetUser(ctx)
	id := ctx.Params("id")
	request := &model.DeleteCategoryRequest{
		UserID: auth.ID,
		ID:     id,
	}
	if err := c.Service.Delete(ctx.UserContext(), request); err != nil {
		c.Log.WithError(err).Error("Failed to delete category")
		return err
	}
	return ctx.JSON(model.WebResponse[bool]{Data: true})
}

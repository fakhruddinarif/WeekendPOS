package controller

import (
	"WeekendPOS/app/delivery/http/middleware"
	"WeekendPOS/app/model"
	"WeekendPOS/app/service"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"math"
	"strconv"
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
	sku := ctx.FormValue("sku")
	name := ctx.FormValue("name")
	categoryID := ctx.FormValue("category_id")
	buyPrice, _ := strconv.ParseFloat(ctx.FormValue("buy_price"), 64)
	sellPrice, _ := strconv.ParseFloat(ctx.FormValue("sell_price"), 64)
	stock, _ := strconv.Atoi(ctx.FormValue("stock"))

	request := &model.CreateProductRequest{
		UserID:     auth.ID,
		SKU:        sku,
		Name:       name,
		CategoryID: categoryID,
		BuyPrice:   buyPrice,
		SellPrice:  sellPrice,
		Stock:      stock,
	}

	fileHeader, err := ctx.FormFile("photo")
	if err != nil {
		c.Log.Warnf("Failed to retrieve file: %+v", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Failed to retrieve file"})
	}

	response, err := c.Service.Create(ctx.UserContext(), request, fileHeader)
	if err != nil {
		c.Log.WithError(err).Error("Failed to create product")
		return err
	}
	return ctx.JSON(model.WebResponse[*model.ProductResponse]{Data: response})
}

func (c *ProductController) Get(ctx *fiber.Ctx) error {
	auth := middleware.GetUser(ctx)

	request := &model.GetProductRequest{
		ID:     ctx.Params("id"),
		UserID: auth.ID,
	}

	response, err := c.Service.Get(ctx.UserContext(), request)
	if err != nil {
		c.Log.WithError(err).Error("Failed to get product")
		return err
	}
	return ctx.JSON(model.WebResponse[*model.ProductResponse]{Data: response})
}

func (c *ProductController) List(ctx *fiber.Ctx) error {
	auth := middleware.GetUser(ctx)

	request := &model.SearchProductRequest{
		UserID: auth.ID,
		Name:   ctx.Query("name", ""),
		Page:   ctx.QueryInt("page", 1),
		Size:   ctx.QueryInt("size", 10),
	}

	response, total, err := c.Service.List(ctx.UserContext(), request)
	if err != nil {
		c.Log.WithError(err).Error("Failed to search product")
		return err
	}
	paging := &model.PageMetadata{
		Page:      request.Page,
		Size:      request.Size,
		TotalItem: total,
		TotalPage: int64(math.Ceil(float64(total) / float64(request.Size))),
	}
	return ctx.JSON(model.WebResponse[[]model.ProductResponse]{Data: response, Paging: paging})
}

func (c *ProductController) Update(ctx *fiber.Ctx) error {
	auth := middleware.GetUser(ctx)

	request := new(model.UpdateProductRequest)
	if err := ctx.BodyParser(request); err != nil {
		c.Log.WithError(err).Error("Failed to parse request body")
		return fiber.ErrBadRequest
	}
	request.UserID = auth.ID
	request.ID = ctx.Params("id")

	response, err := c.Service.Update(ctx.UserContext(), request)
	if err != nil {
		c.Log.WithError(err).Error("Failed to update product")
		return err
	}
	return ctx.JSON(model.WebResponse[*model.ProductResponse]{Data: response})
}

func (c *ProductController) Delete(ctx *fiber.Ctx) error {
	auth := middleware.GetUser(ctx)

	request := &model.DeleteProductRequest{
		ID:     ctx.Params("id"),
		UserID: auth.ID,
	}

	err := c.Service.Delete(ctx.UserContext(), request)
	if err != nil {
		c.Log.WithError(err).Error("Failed to delete product")
		return err
	}
	return ctx.JSON(model.WebResponse[bool]{Data: true})
}

func (c *ProductController) AddStock(ctx *fiber.Ctx) error {
	auth := middleware.GetUser(ctx)

	request := new(model.UpdateProductRequest)
	if err := ctx.BodyParser(request); err != nil {
		c.Log.WithError(err).Error("Failed to parse request body")
		return fiber.ErrBadRequest
	}
	request.UserID = auth.ID

	response, err := c.Service.Update(ctx.UserContext(), request)
	if err != nil {
		c.Log.WithError(err).Error("Failed to add stock product")
		return err
	}
	return ctx.JSON(model.WebResponse[*model.ProductResponse]{Data: response})
}

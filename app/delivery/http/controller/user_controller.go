package controller

import (
	"WeekendPOS/app/delivery/http/middleware"
	"WeekendPOS/app/model"
	"WeekendPOS/app/service"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type UserController struct {
	Log     *logrus.Logger
	Service *service.UserService
}

func NewUserController(service *service.UserService, logger *logrus.Logger) *UserController {
	return &UserController{
		Log:     logger,
		Service: service,
	}
}

func (c *UserController) Register(ctx *fiber.Ctx) error {
	name := ctx.FormValue("name")
	username := ctx.FormValue("username")
	password := ctx.FormValue("password")
	email := ctx.FormValue("email")
	phone := ctx.FormValue("phone")

	request := &model.RegisterUserRequest{
		Name:     name,
		Username: username,
		Password: password,
		Email:    email,
		Phone:    phone,
	}

	fileHeader, err := ctx.FormFile("photo")
	if err != nil {
		c.Log.Warnf("Failed to retrieve file: %+v", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Failed to retrieve file"})
	}

	response, err := c.Service.Create(ctx.UserContext(), request, fileHeader)
	if err != nil {
		c.Log.Warnf("Failed to register user : %+v", err)
		return err
	}

	return ctx.JSON(model.WebResponse[*model.UserResponse]{Data: response})
}

func (c *UserController) Login(ctx *fiber.Ctx) error {
	request := new(model.LoginUserRequest)
	err := ctx.BodyParser(request)
	if err != nil {
		c.Log.Warnf("Failed to parse request body : %+v", err)
		return fiber.ErrBadRequest
	}

	response, err := c.Service.Login(ctx.UserContext(), request)
	if err != nil {
		c.Log.Warnf("Failed to login user : %+v", err)
		return err
	}

	return ctx.JSON(model.WebResponse[*model.UserResponse]{Data: response})
}

func (c *UserController) Get(ctx *fiber.Ctx) error {
	auth := middleware.GetUser(ctx)
	request := &model.GetUserRequest{ID: auth.ID}

	response, err := c.Service.Get(ctx.UserContext(), request)
	if err != nil {
		c.Log.Warnf("Failed to get user : %+v", err)
		return err
	}
	return ctx.JSON(model.WebResponse[*model.UserResponse]{Data: response})
}

func (c *UserController) Update(ctx *fiber.Ctx) error {
	auth := middleware.GetUser(ctx)

	name := ctx.FormValue("name")
	username := ctx.FormValue("username")
	password := ctx.FormValue("password")
	email := ctx.FormValue("email")
	phone := ctx.FormValue("phone")

	request := &model.UpdateUserRequest{
		ID:       auth.ID,
		Name:     name,
		Email:    email,
		Phone:    phone,
		Username: username,
		Password: password,
	}

	fileHeader, err := ctx.FormFile("photo")
	if err != nil {
		c.Log.Warnf("Failed to retrieve file: %+v", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Failed to retrieve file"})
	}

	response, err := c.Service.Update(ctx.UserContext(), request, fileHeader)
	if err != nil {
		c.Log.Warnf("Failed to update user : %+v", err)
		return err
	}

	return ctx.JSON(model.WebResponse[*model.UserResponse]{Data: response})
}

func (c *UserController) Logout(ctx *fiber.Ctx) error {
	auth := middleware.GetUser(ctx)

	request := &model.LogoutUserRequest{ID: auth.ID}

	response, err := c.Service.Logout(ctx.UserContext(), request)
	if err != nil {
		c.Log.Warnf("Failed to logout user : %+v", err)
		return err
	}

	return ctx.JSON(model.WebResponse[bool]{Data: response})
}

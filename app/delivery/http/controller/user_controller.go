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
	request := new(model.RegisterUserRequest)
	err := ctx.BodyParser(request)
	if err != nil {
		c.Log.Warnf("Failed to parse request body : %+v", err)
		return fiber.ErrBadRequest
	}

	response, err := c.Service.Create(ctx.UserContext(), request)
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

	request := new(model.UpdateUserRequest)
	if err := ctx.BodyParser(request); err != nil {
		c.Log.Warnf("Failed to parse request body : %+v", err)
		return fiber.ErrBadRequest
	}

	request.ID = auth.ID
	response, err := c.Service.Update(ctx.UserContext(), request)
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

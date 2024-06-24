package route

import (
	"WeekendPOS/app/delivery/http/controller"
	"github.com/gofiber/fiber/v2"
)

type RouteConfig struct {
	App            *fiber.App
	UserController *controller.UserController
	AuthMiddleware fiber.Handler
}

func (c *RouteConfig) Setup() {
	c.SetupGuestRoute()
	c.SetupAuthRoute()
}

func (c *RouteConfig) SetupGuestRoute() {
	c.App.Post("/api/register", c.UserController.Register)
	c.App.Post("/api/login", c.UserController.Login)
}

func (c *RouteConfig) SetupAuthRoute() {
	c.App.Use(c.AuthMiddleware)
	c.App.Get("/api/user", c.UserController.Get)
	c.App.Patch("/api/user", c.UserController.Update)
	c.App.Delete("/api/user/logout", c.UserController.Logout)
}

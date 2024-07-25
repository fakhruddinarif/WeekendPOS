package route

import (
	"WeekendPOS/app/delivery/http/controller"
	"github.com/gofiber/fiber/v2"
)

type RouteConfig struct {
	App                   *fiber.App
	UserController        *controller.UserController
	CategoryController    *controller.CategoryController
	ProductController     *controller.ProductController
	TransactionController *controller.TransactionController
	AuthMiddleware        fiber.Handler
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

	// User
	userRoutes := c.App.Group("/api/user", c.AuthMiddleware)
	userRoutes.Get("/", c.UserController.Get)
	userRoutes.Patch("/", c.UserController.Update)
	userRoutes.Delete("/", c.UserController.Logout)
	userRoutes.Post("/add_employee", c.UserController.AddEmployee)
	userRoutes.Get("/employees", c.UserController.ListEmployees)

	// Category
	categoryRoutes := c.App.Group("/api/category", c.AuthMiddleware)
	categoryRoutes.Post("/", c.CategoryController.Create)
	categoryRoutes.Get("/", c.CategoryController.List)
	categoryRoutes.Get("/:id", c.CategoryController.Get)
	categoryRoutes.Put("/:id", c.CategoryController.Update)
	categoryRoutes.Delete("/:id", c.CategoryController.Delete)

	// Product
	productRoutes := c.App.Group("/api/product", c.AuthMiddleware)
	productRoutes.Post("/", c.ProductController.Create)
	productRoutes.Get("/", c.ProductController.List)
	productRoutes.Get("/:id", c.ProductController.Get)
	productRoutes.Put("/:id", c.ProductController.Update)
	productRoutes.Delete("/:id", c.ProductController.Delete)
	productRoutes.Patch("/add_stock", c.ProductController.AddStock)

	// Transaction
	transactionRoutes := c.App.Group("/api/transaction", c.AuthMiddleware)
	transactionRoutes.Post("/", c.TransactionController.Create)
}

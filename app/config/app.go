package config

import (
	"WeekendPOS/app/delivery/http/controller"
	"WeekendPOS/app/delivery/http/middleware"
	"WeekendPOS/app/delivery/http/route"
	"WeekendPOS/app/gateway/messaging"
	"WeekendPOS/app/repository"
	"WeekendPOS/app/service"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type BootstrapConfig struct {
	DB       *gorm.DB
	App      *fiber.App
	Log      *logrus.Logger
	Validate *validator.Validate
	Config   *viper.Viper
	Producer *kafka.Producer
}

func Bootstrap(config *BootstrapConfig) {
	// setup repositories
	userRepository := repository.NewUserRepository(config.Log)
	categoryRepository := repository.NewCategoryRepository(config.Log)

	// setup producer
	userProducer := messaging.NewUserProducer(config.Producer, config.Log)
	categoryProducer := messaging.NewCategoryProducer(config.Producer, config.Log)

	// setup service
	userService := service.NewUserService(config.DB, config.Log, config.Validate, userRepository, userProducer)
	categoryService := service.NewCategoryService(config.DB, config.Log, config.Validate, categoryRepository, categoryProducer)

	// setup controller
	userController := controller.NewUserController(userService, config.Log)
	categoryController := controller.NewCategoryController(categoryService, config.Log)

	// setup middleware
	authMiddleware := middleware.NewAuth(userService)

	routeConfig := route.RouteConfig{
		App:                config.App,
		UserController:     userController,
		CategoryController: categoryController,
		AuthMiddleware:     authMiddleware,
	}
	routeConfig.Setup()
}

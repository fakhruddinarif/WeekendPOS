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

	// setup producer
	userProducer := messaging.NewUserProducer(config.Producer, config.Log)

	// setup use cases
	userUseCase := service.NewUserService(config.DB, config.Log, config.Validate, userRepository, userProducer)

	// setup controller
	userController := controller.NewUserController(userUseCase, config.Log)

	// setup middleware
	authMiddleware := middleware.NewAuth(userUseCase)

	routeConfig := route.RouteConfig{
		App:            config.App,
		UserController: userController,
		AuthMiddleware: authMiddleware,
	}
	routeConfig.Setup()
}

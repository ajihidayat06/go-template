package router

import (
	"go-template/config"
	"go-template/controller/bookcontroller.go"
	"go-template/middleware"
	"go-template/usecase"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type InitRouterStruct struct {
	BookController bookcontroller.BookController
}

func InitRouter(initUseCase *usecase.InitUseCaseStruct, cfg *config.CfgStruct) InitRouterStruct {
	return InitRouterStruct{
		BookController: bookcontroller.NewBookController(&initUseCase.BookUseCase, cfg),
	}
}

func (r *InitRouterStruct) SetupRouting(app *fiber.App, logger *logrus.Logger) {
	app.Use(middleware.PanicRecoveryMiddleware())
	app.Use(middleware.LoggingMiddleware(logger))
	r.BookRoute(app)
}

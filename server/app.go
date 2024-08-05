package server

import (
	"go-template/config"
	"go-template/repo"
	"go-template/router"
	"go-template/usecase"
	"log"
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var (
	Validator = validator.New()
)

func AppRun() {
	app := fiber.New()
	logger := config.InitLogger()

	var err error
	config.DB, err = config.ConnectDB()
	if err != nil {
		logger.Error("can't, connect to db")
		os.Exit(1)
	}

	cfg := config.NewCfg()

	repo := repo.InitRepo(config.DB, &cfg)

	usecase := usecase.InitUseCase(&repo, config.DB, &cfg)

	router := router.InitRouter(&usecase, &cfg)
	router.SetupRouting(app, logger)

	err = app.Listen(":8080")
	if err != nil {
		log.Fatal(err)
	}
}

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

	cfg := config.NewCfg()
	config.Validator = validator.New()
	config.Loging = config.InitLogger()

	var err error
	config.DB, err = config.ConnectDB()
	if err != nil {
		config.Loging.LogError("can't, connect to db")
		os.Exit(1)
	}

	repo := repo.InitRepo(config.DB, &cfg)

	usecase := usecase.InitUseCase(&repo, config.DB, &cfg)

	router := router.InitRouter(&usecase, &cfg)
	router.SetupRouting(app, &cfg)

	err = app.Listen(":8080")
	if err != nil {
		log.Fatal(err)
	}
}

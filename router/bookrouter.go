package router

import (
	"github.com/gofiber/fiber/v2"
)

func (r *InitRouterStruct) BookRoute(app *fiber.App) {
	app.Post("/api/book", r.BookController.InsertBook)
}

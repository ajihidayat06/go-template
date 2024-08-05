package model

import "github.com/gofiber/fiber/v2"

type Response struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
	Data       any    `json:"data"`
	ErrDetail  any    `json:"err_detail"`
}

func (r *Response) SetResponseJson(ctx *fiber.Ctx) error {
	return ctx.Status(r.StatusCode).JSON(r)
}

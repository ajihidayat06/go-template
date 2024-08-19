package middleware

import (
	"encoding/json"
	"fmt"
	"go-template/config"
	"go-template/model"
	"time"

	"github.com/gofiber/fiber/v2"
)

func PanicRecoveryMiddleware(cfg *config.CfgStruct) fiber.Handler {
	return func(c *fiber.Ctx) error {
		defer func() {
			if r := recover(); r != nil {
				config.Loging.LogPanic(r)
				c.Status(fiber.StatusInternalServerError).JSON(model.Response{
					StatusCode: fiber.StatusInternalServerError,
					Message:    fiber.ErrInternalServerError.Message,
				})
			}
		}()
		return c.Next()
	}
}

// Middleware untuk mencatat log
func LoggingMiddleware(cfg *config.CfgStruct) fiber.Handler {
	return func(c *fiber.Ctx) error {
		start := time.Now()

		// Process request
		err := c.Next()

		// Stop timer
		stop := time.Since(start)

		// Log request details
		config.Loging.LogMessage = config.LogMessage{
			Method:    c.Method(),
			Path:      c.Path(),
			Status:    c.Response().StatusCode(),
			Latency:   stop,
			ClientIp:  c.IP(),
			UserAgent: c.Get("User-Agent"),
			ErrFile:   c.Locals("err_file"),
		}

		if err != nil {
			config.Loging.LogError(err.Error())
		} else if c.Response().StatusCode() >= 400 {
			var resp model.Response
			_ = json.Unmarshal(c.Response().Body(), &resp)
			config.Loging.LogError(fmt.Errorf(fmt.Sprintf(`%v`, resp.ErrDetail)))
		} else {
			config.Loging.LogInfo("Request processed successfully")
		}

		return err
	}
}

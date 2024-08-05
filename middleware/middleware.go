package middleware

import (
	"encoding/json"
	"go-template/model"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

func PanicRecoveryMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		defer func() {
			if r := recover(); r != nil {
				log.Printf("Recovered from panic: %v", r)
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
func LoggingMiddleware(logger *logrus.Logger) fiber.Handler {
	return func(c *fiber.Ctx) error {
		start := time.Now()

		// Process request
		err := c.Next()

		// Stop timer
		stop := time.Since(start)

		// Log request details
		logEntry := logger.WithFields(logrus.Fields{
			"method":     c.Method(),
			"path":       c.Path(),
			"status":     c.Response().StatusCode(),
			"latency":    stop,
			"client_ip":  c.IP(),
			"user_agent": c.Get("User-Agent"),
		})

		if err != nil {
			logEntry.WithField("error", err.Error()).Error("Server error occurred")
		} else if c.Response().StatusCode() >= 400 {
			var resp model.Response
			_ = json.Unmarshal(c.Response().Body(), &resp)

			if c.Response().StatusCode() >= 500 {
				logEntry.WithField("error", resp.ErrDetail).Error("Server error occurred")
			} else {
				logEntry.WithField("error", resp.ErrDetail).Warn("Client error occurred")
			}
		} else {
			logEntry.Info("Request processed successfully")
		}

		return err
	}
}

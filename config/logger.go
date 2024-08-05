package config

import (
	"os"

	"github.com/sirupsen/logrus"
)

// Konfigurasi logger untuk mencatat ke konsol
func InitLogger() *logrus.Logger {
	logger := logrus.New()
	logger.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
	logger.SetOutput(os.Stdout)
	logger.SetLevel(logrus.InfoLevel) // Sesuaikan level log sesuai kebutuhan
	return logger
}

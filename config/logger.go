package config

import (
	"fmt"
	"go-template/common"
	"os"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
)

const (
	c_DARK_RED = "\033[31m"
	c_YELLOW   = "\033[33m"
	c_GREEN    = "\033[32m"
	c_RESET    = "\033[0m"
)

type LoggerStruct struct {
	Log        *logrus.Logger
	LogMessage LogMessage
}

type LogMessage struct {
	Method     string
	Path       string
	Status     int
	Latency    time.Duration
	ClientIp   string
	UserAgent  string
	StactTrace string
	ErrFile    interface{}
}

// Konfigurasi logger untuk mencatat ke konsol
func InitLogger() *LoggerStruct {
	logger := logrus.New()
	logger.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
		ForceColors:   true,
	})
	logger.SetOutput(os.Stdout)
	logger.SetLevel(logrus.InfoLevel) // Sesuaikan level log sesuai kebutuhan

	return &LoggerStruct{
		Log: logger,
	}
}

func green() string {
	return fmt.Sprint(c_GREEN, "[INFO]", c_RESET, " -- ")
}

func red(msg string) string {
	return fmt.Sprint(c_DARK_RED, msg, c_RESET, " -- ")
}

func (l *LoggerStruct) setMessage(msg ...any) string {
	result := strings.Join(common.ArrInterfaceToArrStr(msg...), " ")

	if common.IsStructPopulated(l.LogMessage) {
		result = fmt.Sprintf("%v \r\n [method:%s | path:%s | status:%d | latency:%s | client_ip:%s | user_agent:%s] \r\n :> %v",
			l.LogMessage.ErrFile, l.LogMessage.Method, l.LogMessage.Path, l.LogMessage.Status, l.LogMessage.Latency, l.LogMessage.ClientIp, l.LogMessage.UserAgent, result)
	}

	return result
}

func (l *LoggerStruct) LogInfo(msg ...any) {
	l.Log.Info(green(), l.setMessage(msg...))
}

func (l *LoggerStruct) LogError(msg ...any) {
	l.Log.Error(red("[ERROR]"), l.setMessage(msg...))
}

func (l *LoggerStruct) LogPanic(msg ...any) {
	l.Log.Log(logrus.ErrorLevel, red("[PANIC]"), l.setMessage(msg...))
}

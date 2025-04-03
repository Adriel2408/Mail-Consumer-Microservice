package logger

import (
	"os"
	"time"

	"github.com/sirupsen/logrus"
)

func NewLogger(logDir string) *logrus.Logger {
	log := logrus.New()
	log.SetFormatter(&logrus.JSONFormatter{})

	logFile := logDir + "/log_" + time.Now().Format("20060102") + ".log"
	file, err := os.OpenFile(logFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic("Failed to open log file: " + err.Error())
	}
	log.SetOutput(file)

	return log
}

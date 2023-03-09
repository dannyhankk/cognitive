package util

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"os"
)

var (
	Logger        *log.Logger
	ContentLogger *log.Logger
)

func LogInit(filePath string) error {
	Logger = log.New()
	logFile, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return fmt.Errorf("open log file error, %s", err.Error())
	}
	Logger.SetReportCaller(true)
	Logger.SetOutput(logFile)
	Logger.SetFormatter(&log.JSONFormatter{})

	ContentLogger = log.New()
	ContentLogFile, err := os.OpenFile("./content.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return fmt.Errorf("open content log file error, %s", err.Error())
	}
	ContentLogger.SetOutput(ContentLogFile)
	return nil
}

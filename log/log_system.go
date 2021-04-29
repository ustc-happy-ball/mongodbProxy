package log

import (
	"github.com/TianqiS/database_for_happyball/configs"
	"github.com/sirupsen/logrus"
	"os"
)

func InitLogSystem() {
	logLevel, err := logrus.ParseLevel(configs.LogLevel)
	if err != nil {
		logrus.SetLevel(logrus.DebugLevel)
		go logrus.Error()
	} else {
		logrus.SetLevel(logLevel)
	}
	if configs.LogToFile {
		var logFile string
		if configs.LogLevel == "info" {
			logFile = configs.LogFilePathProduction
		} else {
			logFile = configs.LogFilePathDevelopment
		}
		f, err := os.OpenFile(logFile, os.O_WRONLY | os.O_CREATE | os.O_APPEND, 0755)
		if err != nil {
			go logrus.Error(err)
		} else {
			logrus.SetOutput(f)
		}
		customFormatter := &logrus.JSONFormatter{}
		logrus.SetFormatter(customFormatter)
	} else {
		customFormatter := &logrus.TextFormatter{}
		customFormatter.FullTimestamp = true
		customFormatter.TimestampFormat = "2006-01-02 15:04:05"
		customFormatter.ForceColors = true
		logrus.SetFormatter(customFormatter)
	}
}
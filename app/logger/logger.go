package logger

import (
	"fmt"
	"io"
	"os"
	"sync"
	"time"

	. "session-service-v2/app/utils"

	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

var lock = &sync.Mutex{}
var logger *FpayLogger
var logFilePath, env, logLevel, serviceName, hostName, country string

func init() {
	godotenv.Load()
	logFilePath = GetStringEnv("LOG_FILE_PATH", "/var/log/pm2/")
	logLevel = GetStringEnv("LOG_LEVEL", "info")
	serviceName = GetStringEnv("SERVICE_NAME", "api-sessions")
	env = GetStringEnv("ENV", "local")
	country = GetStringEnv("COUNTRY", "-")
	hostName, _ = os.Hostname()
}

type FpayLogger struct {
	logEntry *log.Entry
}

func (logger *FpayLogger) Info(message string, args ...interface{}) {
	if len(args) == 0 {
		logger.logEntry.Info(message)
	} else {
		logger.logEntry.WithField("metadata", args[0]).Info(message)
	}
}

func (logger *FpayLogger) Warn(message string, args ...interface{}) {
	if len(args) == 0 {
		logger.logEntry.Warn(message)
	} else {
		logger.logEntry.WithField("metadata", args[0]).Warn(message)
	}
}

func (logger *FpayLogger) Error(message string, args ...interface{}) {
	if len(args) == 0 {
		logger.logEntry.Error(message)
	} else {
		logger.logEntry.WithField("metadata", args[0]).Error(message)
	}
}

func (logger *FpayLogger) Debug(message string, args ...interface{}) {
	if len(args) == 0 {
		logger.logEntry.Debug(message)
	} else {
		logger.logEntry.WithField("metadata", args[0]).Debug(message)
	}
}

func GetLogger() *FpayLogger {
	if logger == nil {
		lock.Lock()
		defer lock.Unlock()
		if logger == nil {
			logger = &FpayLogger{logEntry: createLogger()}
		}
	}

	return logger
}

func createLogger() *log.Entry {

	logger := log.New()

	// Set stdout or file appender
	if env == "local" {
		logger.Out = os.Stdout
	} else {
		logFileName := getFileName(serviceName)
		logFileCompletePath := fmt.Sprintf("%s%s", logFilePath, logFileName)
		file, err := os.OpenFile(logFileCompletePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err == nil {
			logger.Out = io.MultiWriter(os.Stdout, file)
		} else {
			log.Info("Failed to log to file, using default stderr")
		}
	}

	// Set level
	level, _ := log.ParseLevel(logLevel)
	log.SetLevel(level)

	// Formating default fields names
	formatter := &log.JSONFormatter{
		FieldMap: log.FieldMap{
			log.FieldKeyTime:  "@timestamp",
			log.FieldKeyLevel: "level",
			log.FieldKeyMsg:   "description",
			log.FieldKeyFunc:  "caller",
		},
	}
	logger.SetFormatter(formatter)

	// Setting default fields
	return logger.WithFields(log.Fields{
		"country":     country,
		"serviceHost": hostName,
		"service":     serviceName,
		"environment": env,
		"source":      "go",
		"version":     "1.0.0", // TODO: retrieve image version
		"app":         "common",
		"team":        "wallet"})

	// TODO: write header fields

}

// log-{service-name}-18-06-2021-8455fe470e7d.log
func getFileName(serviceName string) string {

	return fmt.Sprintf("log-%s-%s-%s.log",
		serviceName,
		time.Now().Format("01-02-2006"),
		hostName)
}

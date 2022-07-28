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
var logger *log.Entry

var logFilePath, env, logLevel, serviceName string

func init() {
	godotenv.Load()
	logFilePath = GetStringEnv("LOG_FILE_PATH", "/var/log/pm2/")
	logLevel = GetStringEnv("LOG_LEVEL", "info")
	serviceName = GetStringEnv("SERVICE_NAME", "service")
	env = GetStringEnv("ENV", "local")
}

func GetLogger() *log.Entry {
	if logger == nil {
		lock.Lock()
		defer lock.Unlock()
		if logger == nil {
			logger = createLogger()
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
		"service":     serviceName,
		"environment": env})

}

// Util to add default to getenv function
func getenv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}

// log-{service-name}-18-06-2021-8455fe470e7d.log
func getFileName(serviceName string) string {

	hostname, _ := os.Hostname()

	return fmt.Sprintf("log-%s-%s-%s.log",
		serviceName,
		time.Now().Format("01-02-2006"),
		hostname)
}

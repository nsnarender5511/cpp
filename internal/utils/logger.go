package utils

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/sirupsen/logrus"
)

// Logger instance used throughout the application
var log *logrus.Logger

// InitLogger initializes the logging system
func InitLogger(appPaths AppPaths) {
	// Create new logger instance
	log = logrus.New()

	// Set log level from environment
	level := logrus.InfoLevel
	if levelStr := os.Getenv("LOG_LEVEL"); levelStr != "" {
		if lvl, err := logrus.ParseLevel(levelStr); err == nil {
			level = lvl
		}
	}
	log.SetLevel(level)

	// Get log file name from environment or use default
	logFileName := os.Getenv("LOG_FILE_NAME")
	if logFileName == "" {
		logFileName = DefaultLogFileName
	}

	// Ensure log directory exists
	logFilePath := appPaths.GetLogFile(logFileName)
	logDir := filepath.Dir(logFilePath)

	if err := EnsureDirExists(logDir, 0755); err == nil {
		// Set up file for logging
		file, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
		if err == nil {
			// Create multi-writer to log to both file and stdout
			mw := io.MultiWriter(os.Stdout, file)
			log.SetOutput(mw)
		}
	}

	// Use a simpler text formatter
	log.SetFormatter(&logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
	})
}

// Debug logs a debug message
func Debug(msg string) {
	log.Debug(msg)
}

// Info logs an informational message
func Info(msg string) {
	log.Info(msg)
}

// Warn logs a warning message
func Warn(msg string) {
	log.Warn(msg)
}

// Error logs an error message
func Error(msg string) {
	log.Error(msg)
}

// Fatal logs a fatal message and exits
func Fatal(msg string) {
	log.Fatal(msg)
}

// Debugf logs a formatted debug message
func Debugf(format string, args ...interface{}) {
	log.Debug(fmt.Sprintf(format, args...))
}

// Infof logs a formatted informational message
func Infof(format string, args ...interface{}) {
	log.Info(fmt.Sprintf(format, args...))
}

// Warnf logs a formatted warning message
func Warnf(format string, args ...interface{}) {
	log.Warn(fmt.Sprintf(format, args...))
}

// Errorf logs a formatted error message
func Errorf(format string, args ...interface{}) {
	log.Error(fmt.Sprintf(format, args...))
}

// Fatalf logs a formatted fatal message and exits
func Fatalf(format string, args ...interface{}) {
	log.Fatal(fmt.Sprintf(format, args...))
}

package utils

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/sirupsen/logrus"
)

// Logger instance used throughout the application
var log *logrus.Logger

// Console output control flags
var (
	debugConsole bool
	verbose      bool
)

// SetDebugConsole controls debug message console output
func SetDebugConsole(enabled bool) {
	debugConsole = enabled
}

// SetVerbose controls info message console output
func SetVerbose(enabled bool) {
	verbose = enabled
}

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
			// Only log to file, not to stdout
			log.SetOutput(file)
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

	if debugConsole {
		fmt.Println("[DEBUG]", msg)
	}
}

// Info logs an informational message
func Info(msg string) {
	log.Info(msg)

	if verbose {
		fmt.Println("[INFO]", msg)
	}
}

// Warn logs a warning message
func Warn(msg string) {
	log.Warn(msg)

	// Always show warnings on console
	fmt.Println("[WARN]", msg)
}

// Error logs an error message
func Error(msg string) {
	log.Error(msg)

	// Always show errors on console
	fmt.Println("[ERROR]", msg)
}

// Fatal logs a fatal message and exits
func Fatal(msg string) {
	log.Fatal(msg) // This will call os.Exit(1)
}

// Debugf logs a formatted debug message
func Debugf(format string, args ...interface{}) {
	formatted := fmt.Sprintf(format, args...)
	Debug(formatted)
}

// Infof logs a formatted informational message
func Infof(format string, args ...interface{}) {
	formatted := fmt.Sprintf(format, args...)
	Info(formatted)
}

// Warnf logs a formatted warning message
func Warnf(format string, args ...interface{}) {
	formatted := fmt.Sprintf(format, args...)
	Warn(formatted)
}

// Errorf logs a formatted error message
func Errorf(format string, args ...interface{}) {
	formatted := fmt.Sprintf(format, args...)
	Error(formatted)
}

// Fatalf logs a formatted fatal message and exits
func Fatalf(format string, args ...interface{}) {
	formatted := fmt.Sprintf(format, args...)
	Fatal(formatted)
}

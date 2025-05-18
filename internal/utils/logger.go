package utils

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/sirupsen/logrus"
)


var log *logrus.Logger


var (
	debugConsole  bool
	verbose       bool
	verboseErrors bool
)


func SetDebugConsole(enabled bool) {
	debugConsole = enabled
}


func SetVerbose(enabled bool) {
	verbose = enabled
}


func SetVerboseErrors(enabled bool) {
	verboseErrors = enabled
}


func IsVerboseErrors() bool {
	return verboseErrors
}


func IsDebug() bool {
	return debugConsole
}


func IsVerbose() bool {
	return verbose
}


func InitLogger(appPaths AppPaths) {
	
	log = logrus.New()

	
	level := logrus.InfoLevel
	if levelStr := os.Getenv("LOG_LEVEL"); levelStr != "" {
		if lvl, err := logrus.ParseLevel(levelStr); err == nil {
			level = lvl
		}
	}
	log.SetLevel(level)

	
	logFileName := os.Getenv("LOG_FILE_NAME")
	if logFileName == "" {
		logFileName = DefaultLogFileName
	}

	
	logFilePath := appPaths.GetLogFile(logFileName)
	logDir := filepath.Dir(logFilePath)

	if err := EnsureDirExists(logDir, 0755); err == nil {
		
		file, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
		if err == nil {
			
			log.SetOutput(file)
		}
	}

	
	log.SetFormatter(&logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
	})
}


func Debug(msg string) {
	log.Debug(msg)

	if debugConsole {
		fmt.Printf("\033[38;5;246m[DEBUG]\033[0m %s\n", msg)
	}
}


func Info(msg string) {
	log.Info(msg)

	if verbose {
		fmt.Printf("\033[38;5;117m[INFO]\033[0m %s\n", msg)
	}
}


func Warn(msg string) {
	log.Warn(msg)

	
	fmt.Printf("\033[38;5;220m[WARN]\033[0m %s\n", msg)
}


func Error(msg string) {
	log.Error(msg)

	
	fmt.Printf("\033[38;5;196m[ERROR]\033[0m %s\n", msg)
}


func Fatal(msg string) {
	log.Fatal(msg) 
}


func Debugf(format string, args ...interface{}) {
	formatted := fmt.Sprintf(format, args...)
	Debug(formatted)
}


func Infof(format string, args ...interface{}) {
	formatted := fmt.Sprintf(format, args...)
	Info(formatted)
}


func Warnf(format string, args ...interface{}) {
	formatted := fmt.Sprintf(format, args...)
	Warn(formatted)
}


func Errorf(format string, args ...interface{}) {
	formatted := fmt.Sprintf(format, args...)
	Error(formatted)
}


func Fatalf(format string, args ...interface{}) {
	formatted := fmt.Sprintf(format, args...)
	Fatal(formatted)
}

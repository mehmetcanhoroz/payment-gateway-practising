package logger

import (
	"log"
	"os"
)

// Package-level logger instances
var (
	infoLogger  *log.Logger
	debugLogger *log.Logger
	errorLogger *log.Logger
)

// init function to initialize loggers
func init() {
	infoLogger = log.New(os.Stdout, "[INFO] - ", log.Ldate|log.Ltime|log.Lshortfile)
	debugLogger = log.New(os.Stdout, "[DEBUG] - ", log.Ldate|log.Ltime|log.Lshortfile)
	errorLogger = log.New(os.Stderr, "[ERROR] - ", log.Ldate|log.Ltime|log.Lshortfile)
}

// Info logs informational messages
func Info(v ...interface{}) {
	infoLogger.Println(v...)
}

// Debug logs debugging messages
func Debug(v ...interface{}) {
	debugLogger.Println(v...)
}

// Error logs error messages
func Error(v ...interface{}) {
	errorLogger.Println(v...)
}

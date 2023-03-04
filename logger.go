package chatGPT

import (
	"log"
)

// Logger is an interface that defines logging methods
type Logger interface {
	Debug(msg string, val ...interface{}) error
	Info(msg string, val ...interface{}) error
	Warn(msg string, val ...interface{}) error
	Error(msg string, val ...interface{}) error
}

// GLog is a global variable that holds a DefaultLog instance
var GLog *DefaultLog

// DefaultLog is a struct that holds a pointer to a log.Logger instance
type DefaultLog struct {
	Log *log.Logger
}

// Debug logs a debug message with optional values
func (l *DefaultLog) Debug(msg string, val ...interface{}) error {
	l.Log.Printf(msg, val...)
	return nil
}

// Info logs an informational message with optional values
func (l *DefaultLog) Info(msg string, val ...interface{}) error {
	l.Log.Printf(msg, val...)
	return nil
}

// Warn logs a warning message with optional values
func (l *DefaultLog) Warn(msg string, val ...interface{}) error {
	l.Log.Printf(msg, val...)
	return nil
}

// Error logs an error message with optional values and returns an error
func (l *DefaultLog) Error(msg string, val ...interface{}) error {
	l.Log.Fatalf(msg, val...)
	return nil
}

// GetLogger returns a pointer to a DefaultLog instance
// If GLog is nil, a new DefaultLog instance is created and returned
func GetLogger() *DefaultLog {
	if GLog == nil {
		GLog = &DefaultLog{Log: log.Default()}
	}
	return GLog
}

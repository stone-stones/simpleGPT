package chatGPT

import (
	"log"
)

// Logger is an interface that defines logging methods
type Logger interface {
	Debug(val ...interface{}) error
	Info(val ...interface{}) error
	Warn( val ...interface{}) error
	Error(val ...interface{}) error
}

// GLog is a global variable that holds a DefaultLog instance
var GLog *DefaultLog

// DefaultLog is a struct that holds a pointer to a log.Logger instance
type DefaultLog struct {
	Log *log.Logger
}

func getMsg(val ...interface{})(string ,[]interface{}) {
	if len(val) <= 0  {
		return "",[]interface{}{}
	}
	msg ,ok := val[0].(string)
	if !ok {
		return "",val
	}
	return msg,val[1:]

}
// Debug logs a debug message with optional values
func (l *DefaultLog) Debug(val ...interface{}) error {
	msg,params := getMsg(val...)
	l.Log.Printf(msg, params...)
	return nil
}

// Info logs an informational message with optional values
func (l *DefaultLog) Info(val ...interface{}) error {
	msg,params := getMsg(val...)
	l.Log.Printf(msg,params...)
	return nil
}

// Warn logs a warning message with optional values
func (l *DefaultLog) Warn( val ...interface{}) error {
	msg,params := getMsg(val...)
	l.Log.Printf(msg,params...)
	return nil
}

// Error logs an error message with optional values and returns an error
func (l *DefaultLog) Error(val ...interface{}) error {
	msg,params := getMsg(val...)
	l.Log.Fatalf(msg, params...)
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

package chatGPT

import (
	"log"
)

// Logger is an interface that defines logging methods
type Logger interface {
	Debug(val ...interface{})
	Info(val ...interface{})
	Warn( val ...interface{})
	Error(val ...interface{})
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
func (l *DefaultLog) Debug(val ...interface{})  {
	msg,params := getMsg(val...)
	l.Log.Printf(msg, params...)
}

// Info logs an informational message with optional values
func (l *DefaultLog) Info(val ...interface{})  {
	msg,params := getMsg(val...)
	l.Log.Printf(msg,params...)
}

// Warn logs a warning message with optional values
func (l *DefaultLog) Warn( val ...interface{})  {
	msg,params := getMsg(val...)
	l.Log.Printf(msg,params...)
}

// Error logs an error message with optional values and returns an error
func (l *DefaultLog) Error(val ...interface{})  {
	msg,params := getMsg(val...)
	l.Log.Fatalf(msg, params...)
}

// GetLogger returns a pointer to a DefaultLog instance
// If GLog is nil, a new DefaultLog instance is created and returned
func GetLogger() *DefaultLog {
	if GLog == nil {
		GLog = &DefaultLog{Log: log.Default()}
	}
	return GLog
}

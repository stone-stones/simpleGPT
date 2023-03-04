package chatGPT

import (
	"log"
)
type Logger interface {
	Debug(msg string, val ...interface{}) error
	Info(msg string, val ...interface{}) error
	Warn(msg string, val ...interface{}) error
	Error(msg string, val ...interface{}) error
}

var GLog *DefaultLog
type DefaultLog struct {
	Log *log.Logger
}

func (l *DefaultLog)Debug(msg string ,val ...interface{}) error  {
	l.Log.Printf(msg,val...)
	return nil
}
func (l *DefaultLog)Info(msg string ,val ...interface{}) error  {
	l.Log.Printf(msg,val...)
	return nil
}
func (l *DefaultLog)Warn(msg string ,val ...interface{}) error  {
	l.Log.Printf(msg,val...)
	return nil
}
func (l *DefaultLog)Error(msg string ,val ...interface{}) error  {
	l.Log.Fatalf(msg,val...)
	return nil
}

func GetLogger()*DefaultLog {
	if GLog == nil {
		GLog = &DefaultLog{Log:log.Default()}
	}
	return GLog
}
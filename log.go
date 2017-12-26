package golog

import (
	"log"
	"os"
)

const (
	InfoHeader    = "[\033[36m*\033[0m] "
	WarningHeader = "[\033[33m!\033[0m] "
	DebugHeader   = "[\033[31mDEBUG\033[0m] "
	ErrorHeader   = "[\033[41;37mERROR\033[0m] "
	SuccessHeader = "[\033[32m+\033[0m] "
	FailureHeader = "[\033[31m-\033[0m] "
	PanicHeader   = "[\033[41;37mPanic\033[0m] "
)

type Logger struct {
	*log.Logger
	hasColor bool
}

var Std = NewLogger(log.New(os.Stderr, "", log.LstdFlags))

func NewLogger(l *log.Logger) *Logger {
	return &Logger{l, true}
}

func (l *Logger) Info(v ...interface{}) {
	l.println(InfoHeader, v)
}

func (l *Logger) Infof(format string, v ...interface{}) {
	l.printf(InfoHeader, format, v...)
}

func (l *Logger) Warning(v ...interface{}) {
	l.println(WarningHeader, v)
}

func (l *Logger) Warningf(format string, v ...interface{}) {
	l.printf(WarningHeader, format, v...)
}

func (l *Logger) Debug(v ...interface{}) {
	l.println(DebugHeader, v)
}

func (l *Logger) Debugf(format string, v ...interface{}) {
	l.printf(DebugHeader, format, v...)
}

func (l *Logger) Error(v ...interface{}) {
	l.println(ErrorHeader, v)
}

func (l *Logger) Errorf(format string, v ...interface{}) {
	l.printf(ErrorHeader, format, v...)
}

func (l *Logger) Success(v ...interface{}) {
	l.println(SuccessHeader, v)
}

func (l *Logger) Successf(format string, v ...interface{}) {
	l.printf(SuccessHeader, format, v...)
}
func (l *Logger) Failure(v ...interface{}) {
	l.println(FailureHeader, v)
}

func (l *Logger) Failuref(format string, v ...interface{}) {
	l.printf(FailureHeader, format, v...)
}

func (l *Logger) Panic(v ...interface{}) {
	if l.hasColor {
		oldPrefix := l.Prefix()
		l.SetPrefix(PanicHeader + oldPrefix)
		defer l.SetPrefix(oldPrefix)
	}
	l.Logger.Panic(v...)
}

func (l *Logger) Panicf(format string, v ...interface{}) {
	if l.hasColor {
		oldPrefix := l.Prefix()
		l.SetPrefix(PanicHeader + oldPrefix)
		defer l.SetPrefix(oldPrefix)
	}
	l.Logger.Panicf(format, v...)

}

func (l *Logger) println(header string, v ...interface{}) {
	if l.hasColor {
		oldPrefix := l.Prefix()
		l.SetPrefix(header + oldPrefix)
		defer l.SetPrefix(oldPrefix)
	}
	l.Logger.Println(v...)
}

func (l *Logger) printf(header, format string, v ...interface{}) {
	if l.hasColor {
		oldPrefix := l.Prefix()
		l.SetPrefix(header + oldPrefix)
		defer l.SetPrefix(oldPrefix)
	}
	l.Logger.Printf(format, v...)
}

func (l *Logger) SetColor(b bool) {
	l.ColorFlag = b
}

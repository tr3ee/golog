package golog

import (
	"log"
	"os"
	"testing"
)

func TestLog(t *testing.T) {
	line := "你好"
	l := log.New(os.Stdout, "", 0)
	logger := NewLogger(l)
	logger.Error(line)
	logger.Info(line)
	logger.Debug(line)
	logger.Warning(line)
	logger.Success(line)
	defer func() {
		if err := recover(); err != nil {
			logger.Failure(err.(string))
		}
	}()
	logger.Panic(line)

}

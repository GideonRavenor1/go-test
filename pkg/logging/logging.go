package logging

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"go_test/pkg/utils"
	"io"
	"os"
	"path"
	"runtime"
)

type writeHook struct {
	Writer    []io.Writer
	LogLevels []logrus.Level
}

func (hook *writeHook) Fire(entry *logrus.Entry) error {
	line, err := entry.String()
	if err != nil {
		return err
	}
	for _, w := range hook.Writer {
		_, err := w.Write([]byte(line))
		utils.ErrorHandler(err)
	}
	return err
}

func (hook *writeHook) Levels() []logrus.Level {
	return hook.LogLevels
}

var e *logrus.Entry

type Logger struct {
	*logrus.Entry
}

func GetLogger() Logger {
	return Logger{e}
}

func (logger *Logger) GetLoggerWithField(k string, v interface{}) Logger {
	return Logger{logger.WithField(k, v)}
}

func init() {
	logger := logrus.New()
	logger.SetReportCaller(true)
	logger.Formatter = &logrus.TextFormatter{
		CallerPrettyfier: func(frame *runtime.Frame) (function string, file string) {
			filename := path.Base(frame.File)
			return fmt.Sprintf("%s()", frame.Function), fmt.Sprintf("%s:%d", filename, frame.Line)
		},
		DisableColors: false,
		FullTimestamp: true,
	}
	dirName := "logs"
	if _, err := os.Stat(dirName); err != nil {
		err := os.MkdirAll(dirName, 0644)
		utils.ErrorHandler(err)
	}
	allFiles, err := os.OpenFile(dirName+"/all.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0640)
	utils.ErrorHandler(err)

	logger.SetOutput(io.Discard)

	logger.AddHook(&writeHook{
		Writer:    []io.Writer{allFiles, os.Stdout},
		LogLevels: logrus.AllLevels,
	})
	logger.SetLevel(logrus.TraceLevel)

	e = logrus.NewEntry(logger)
}

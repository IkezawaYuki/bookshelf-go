package logger

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"strings"
)

var (
	Log *logrus.Logger
)

func init() {
	level, err := logrus.ParseLevel("DEBUG")
	if err != nil {
		level = logrus.DebugLevel
	}
	Log = &logrus.Logger{
		Out:       os.Stdout,
		Formatter: &logrus.JSONFormatter{},
		Level:     level,
	}
}

func Info(msg string, tags ...string) {
	if Log.Level < logrus.InfoLevel {
		return
	}
	Log.WithFields(parseFields(tags...)).Info(msg)
}

func Debug(msg string, tags ...string) {
	if Log.Level < logrus.DebugLevel {
		return
	}
	Log.WithFields(parseFields(tags...)).Debug(msg)
}

func Error(msg string, err error, tags ...string) {
	if Log.Level < logrus.ErrorLevel {
		return
	}
	msg = fmt.Sprintf("%s - ERROR - %v", msg, err)
	Log.WithFields(parseFields(tags...)).Error(msg)
}

func parseFields(tags ...string) logrus.Fields {
	result := make(logrus.Fields, len(tags))
	for _, tags := range tags {
		els := strings.Split(tags, ":")
		result[strings.TrimSpace(els[0])] = strings.TrimSpace(els[1])
	}
	return result
}

package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

type Logger interface {
	Debug(args interface{})
	Info(args interface{})
	Warn(args interface{})
	Error(args interface{})
	Fatal(args interface{})
}

type apiLogger struct {
	log *logrus.Logger
}

func NewApiLogger() *apiLogger {
	var l = logrus.New()
	l.Out = os.Stdout
	l.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
	return &apiLogger{
		log: l,
	}
}

func (l *apiLogger) Debug(args interface{}) {
	l.log.WithFields(logrus.Fields{"event": "general"}).Debug(args)
}

func (l *apiLogger) Info(args interface{}) {
	l.log.WithFields(logrus.Fields{"event": "general"}).Info(args)
}

func (l *apiLogger) Warn(args interface{}) {
	l.log.WithFields(logrus.Fields{"event": "general"}).Warn(args)
}

func (l *apiLogger) Error(args interface{}) {
	l.log.WithFields(logrus.Fields{"event": "general"}).Error(args)
}

func (l *apiLogger) Fatal(args interface{}) {
	l.log.WithFields(logrus.Fields{"event": "general"}).Fatal(args)
}

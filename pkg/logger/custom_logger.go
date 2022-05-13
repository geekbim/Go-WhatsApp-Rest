package logger

import (
	"os"
	"time"

	"github.com/sirupsen/logrus"
)

type LogResponseAPI struct {
	Event         string
	TransactionId string
	StatusCode    int
	ResponseTime  time.Duration
	Method        string
	Request       interface{}
	URL           string
	Message       string
	Response      interface{}
	Err           error
}

type LogEventService struct {
	Event        string
	From         string
	To           string
	Payload      interface{}
	ResponseTime time.Duration
	Message      string
}

func CreateLogResponse(data LogResponseAPI) {
	var log = logrus.New()
	log.Out = os.Stdout
	log.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})

	log.WithFields(logrus.Fields{
		"event":         data.Event,
		"transactionid": data.TransactionId,
		"status_code":   data.StatusCode,
		"response_time": data.ResponseTime,
		"method":        data.Method,
		"request":       data.Request,
		"url":           data.URL,
		"response":      data.Response,
	}).Info(data.Message)
}

func CreateFatalErrorLogResponse(data LogResponseAPI) {
	var log = logrus.New()
	log.Out = os.Stdout
	log.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})

	log.WithFields(logrus.Fields{
		"event":         data.Event,
		"transactionid": data.TransactionId,
		"response_time": data.ResponseTime,
		"method":        data.Method,
		"request":       data.Request,
		"url":           data.URL,
		"err":           data.Err.Error(),
	}).Fatal(data.Message)

}

func CreateErrorLogResponse(data LogResponseAPI) {
	var log = logrus.New()
	log.Out = os.Stdout
	log.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})

	log.WithFields(logrus.Fields{
		"event":         data.Event,
		"transactionid": data.TransactionId,
		"response_time": data.ResponseTime,
		"method":        data.Method,
		"request":       data.Request,
		"url":           data.URL,
		"err":           data.Err.Error(),
	}).Error(data.Message)

}

func CreateLogEventService(data LogEventService) {
	var log = logrus.New()
	log.Out = os.Stdout
	log.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})

	log.WithFields(logrus.Fields{
		"event":         data.Event,
		"from":          data.From,
		"to":            data.To,
		"payload":       data.Payload,
		"response_time": data.ResponseTime,
	}).Info(data.Message)
}

func CreateErrorLogEventService(data LogEventService) {
	var log = logrus.New()
	log.Out = os.Stdout
	log.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})

	log.WithFields(logrus.Fields{
		"event":         data.Event,
		"from":          data.From,
		"to":            data.To,
		"payload":       data.Payload,
		"response_time": data.ResponseTime,
	}).Error(data.Message)
}

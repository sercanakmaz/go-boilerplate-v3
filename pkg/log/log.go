package log

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"time"
)

type loggerKeyType string

const correlationIdKey loggerKeyType = "loggerWithCorrelation"
const WarnLevel = logrus.WarnLevel
const InfoLevel = logrus.InfoLevel

type Field struct {
	Url            string
	HostName       string
	HttpStatusCode int
	Duration       int64
	RequestBody    string
	ResponseBody   string
	HttpMethod     string
	Message        string
	Extra          map[string]interface{}
}

type Logger interface {
	Info(ctx context.Context, message string)
	Warn(ctx context.Context, message string)
	Exception(ctx context.Context, message string, error error)
	RequestResponse(ctx context.Context, withFields Field)
	WithCorrelationId(ctx context.Context, id string) context.Context
	Fatal(ctx context.Context, message string, error error)
	Request(ctx context.Context, withFields Field)
	Response(ctx context.Context, withFields Field)
	ResponseWithLevel(ctx context.Context, withFields Field, level logrus.Level)
	InfoWithExtra(ctx context.Context, message string, dictionary map[string]interface{})
}

type logger struct {
	logRus   *logrus.Entry
	logLevel logrus.Level
}

func (l *logger) InfoWithExtra(ctx context.Context, message string, dictionary map[string]interface{}) {

	var fields = logrus.Fields{}
	for key, value := range dictionary {
		fields[key] = value
	}

	l.withContext(ctx).WithFields(fields).Info(message)
}

func (l *logger) Info(ctx context.Context, message string) {

	l.withContext(ctx).WithFields(logrus.Fields{
		"DateTime": time.Now()}).Info(message)
}

func (l *logger) Warn(ctx context.Context, message string) {

	l.withContext(ctx).WithFields(logrus.Fields{
		"DateTime": time.Now()}).Warn(message)
}

func (l *logger) Fatal(ctx context.Context, message string, error error) {

	l.withContext(ctx).WithFields(logrus.Fields{
		"DateTime":  time.Now(),
		"Exception": error}).Error(message)
}

func (l *logger) Exception(ctx context.Context, message string, error error) {

	l.withContext(ctx).WithFields(logrus.Fields{
		"DateTime":  time.Now(),
		"Exception": error}).Error(message)
}

func (l *logger) RequestResponse(ctx context.Context, withFields Field) {

	var fields = logrus.Fields{
		"DateTime":       time.Now(),
		"RequestBody":    withFields.RequestBody,
		"ResponseBody":   withFields.ResponseBody,
		"HttpMethod":     withFields.HttpMethod,
		"HttpStatusCode": withFields.HttpStatusCode,
		"Duration":       withFields.Duration,
		"HostName":       withFields.HostName,
		"Url":            withFields.Url,
	}

	for key, value := range withFields.Extra {
		fields[key] = value
	}

	l.withContext(ctx).WithFields(fields).Info(withFields.Message)

}

func (l *logger) Request(ctx context.Context, withFields Field) {

	var fields = logrus.Fields{
		"DateTime":       time.Now(),
		"RequestBody":    withFields.RequestBody,
		"ResponseBody":   "",
		"HttpMethod":     withFields.HttpMethod,
		"HttpStatusCode": 102,
		"Duration":       0,
		"HostName":       withFields.HostName,
		"Url":            withFields.Url,
	}

	for key, value := range withFields.Extra {
		fields[key] = value
	}

	l.withContext(ctx).WithFields(fields).Info(withFields.Message)

}

func (l *logger) Response(ctx context.Context, withFields Field) {

	var fields = logrus.Fields{
		"DateTime":       time.Now(),
		"RequestBody":    withFields.RequestBody,
		"ResponseBody":   withFields.ResponseBody,
		"HttpMethod":     withFields.HttpMethod,
		"HttpStatusCode": withFields.HttpStatusCode,
		"Duration":       withFields.Duration,
		"HostName":       withFields.HostName,
		"Url":            withFields.Url,
	}

	for key, value := range withFields.Extra {
		fields[key] = value
	}

	l.withContext(ctx).WithFields(fields).Info(withFields.Message)

}

func (l *logger) ResponseWithLevel(ctx context.Context, withFields Field, level logrus.Level) {

	var fields = logrus.Fields{
		"DateTime":       time.Now(),
		"RequestBody":    withFields.RequestBody,
		"ResponseBody":   withFields.ResponseBody,
		"HttpMethod":     withFields.HttpMethod,
		"HttpStatusCode": withFields.HttpStatusCode,
		"Duration":       withFields.Duration,
		"HostName":       withFields.HostName,
		"Url":            withFields.Url,
	}

	for key, value := range withFields.Extra {
		fields[key] = value
	}

	l.withContext(ctx).WithFields(fields).Logln(level, withFields.Message)

}

func NewLogger() Logger {
	var log = logrus.New()
	log.SetFormatter(new(jsonFormatter))
	log.SetLevel(InfoLevel)
	return &logger{logRus: logrus.NewEntry(log), logLevel: InfoLevel}
}
func NewLoggerWithLevel(level logrus.Level) Logger {
	var log = logrus.New()
	log.SetFormatter(new(jsonFormatter))
	return &logger{logRus: logrus.NewEntry(log), logLevel: level}
}

func (l *logger) withContext(ctx context.Context) *logrus.Entry {
	logger := ctx.Value(correlationIdKey)
	if logger == nil {
		return l.logRus
	}
	var logEntry = (logger.(*logrus.Entry))
	logEntry.Logger.SetLevel(l.logLevel)

	return logEntry
}

func (l *logger) WithCorrelationId(ctx context.Context, id string) context.Context {
	return context.WithValue(ctx, correlationIdKey, l.withContext(ctx).WithFields(logrus.Fields{"CorrelationId": id}))
}

type jsonFormatter struct {
}

func (f *jsonFormatter) Format(entry *logrus.Entry) ([]byte, error) {

	entry.Data["Message"] = entry.Message
	entry.Data["Level"] = entry.Level

	if _, ok := entry.Data["Exception"]; ok {
		entry.Data["Exception"] = fmt.Sprint(entry.Data["Exception"])

	}
	serialized, err := json.Marshal(entry.Data)

	if err != nil {
		return nil, fmt.Errorf("Failed to marshal fields to JSON, %v", err)
	}

	return append(serialized, '\n'), nil
}

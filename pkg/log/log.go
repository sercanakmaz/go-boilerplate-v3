package log

import (
	"context"
	"log/slog"
	"os"
)

type loggerKeyType string

const correlationIdKey loggerKeyType = "loggerWithCorrelation"

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
	Exception(ctx context.Context, message string, err error)
	RequestResponse(ctx context.Context, withFields Field)
	WithCorrelationId(ctx context.Context, id string) context.Context
	Fatal(ctx context.Context, message string, err error)
	Request(ctx context.Context, withFields Field)
	Response(ctx context.Context, withFields Field)
	ResponseWithLevel(ctx context.Context, withFields Field, level slog.Level)
	InfoWithExtra(ctx context.Context, message string, dictionary map[string]interface{})
}

type logger struct {
	slogger *slog.Logger
}

func fieldAttrs(f Field) []any {
	attrs := []any{
		slog.String("RequestBody", f.RequestBody),
		slog.String("ResponseBody", f.ResponseBody),
		slog.String("HttpMethod", f.HttpMethod),
		slog.Int("HttpStatusCode", f.HttpStatusCode),
		slog.Int64("Duration", f.Duration),
		slog.String("HostName", f.HostName),
		slog.String("Url", f.Url),
	}
	for k, v := range f.Extra {
		attrs = append(attrs, slog.Any(k, v))
	}
	return attrs
}

func (l *logger) withContext(ctx context.Context) *slog.Logger {
	val := ctx.Value(correlationIdKey)
	if val == nil {
		return l.slogger
	}
	return val.(*slog.Logger)
}

func (l *logger) InfoWithExtra(ctx context.Context, message string, dictionary map[string]interface{}) {
	attrs := make([]any, 0, len(dictionary)*2)
	for k, v := range dictionary {
		attrs = append(attrs, slog.Any(k, v))
	}
	l.withContext(ctx).InfoContext(ctx, message, attrs...)
}

func (l *logger) Info(ctx context.Context, message string) {
	l.withContext(ctx).InfoContext(ctx, message)
}

func (l *logger) Warn(ctx context.Context, message string) {
	l.withContext(ctx).WarnContext(ctx, message)
}

func (l *logger) Fatal(ctx context.Context, message string, err error) {
	l.withContext(ctx).ErrorContext(ctx, message, slog.Any("Exception", err))
}

func (l *logger) Exception(ctx context.Context, message string, err error) {
	l.withContext(ctx).ErrorContext(ctx, message, slog.Any("Exception", err))
}

func (l *logger) RequestResponse(ctx context.Context, withFields Field) {
	l.withContext(ctx).InfoContext(ctx, withFields.Message, fieldAttrs(withFields)...)
}

func (l *logger) Request(ctx context.Context, withFields Field) {
	attrs := fieldAttrs(Field{
		RequestBody:    withFields.RequestBody,
		ResponseBody:   "",
		HttpMethod:     withFields.HttpMethod,
		HttpStatusCode: 102,
		Duration:       0,
		HostName:       withFields.HostName,
		Url:            withFields.Url,
		Extra:          withFields.Extra,
	})
	l.withContext(ctx).InfoContext(ctx, withFields.Message, attrs...)
}

func (l *logger) Response(ctx context.Context, withFields Field) {
	l.withContext(ctx).InfoContext(ctx, withFields.Message, fieldAttrs(withFields)...)
}

func (l *logger) ResponseWithLevel(ctx context.Context, withFields Field, level slog.Level) {
	l.withContext(ctx).Log(ctx, level, withFields.Message, fieldAttrs(withFields)...)
}

func (l *logger) WithCorrelationId(ctx context.Context, id string) context.Context {
	enriched := l.withContext(ctx).With(slog.String("CorrelationId", id))
	return context.WithValue(ctx, correlationIdKey, enriched)
}

func NewLogger() Logger {
	handler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo})
	return &logger{slogger: slog.New(handler)}
}

func NewLoggerWithLevel(level slog.Level) Logger {
	handler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: level})
	return &logger{slogger: slog.New(handler)}
}

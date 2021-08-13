package log

import (
	"context"
	"os"
	"time"

	"go.uber.org/zap/zapcore"

	"go.uber.org/zap"
)

type key string

const (
	// TraceIDKey 追踪id
	TraceIDKey key = "trace_id"
)

var (
	logger *zap.Logger
)

// TimeFormat 时间戳格式
func TimeFormat(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05"))
}

// MSecondsDurationEncoder 毫秒格式
func MSecondsDurationEncoder(d time.Duration, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendInt64(d.Milliseconds())
}

func init() {
	conf := zap.NewProductionConfig()
	logLevel := os.Getenv("LogDebug")
	if logLevel != "" {
		conf.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
	}
	encoderConf := zapcore.EncoderConfig{
		TimeKey:        "timestamp",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "file",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeTime:     TimeFormat,
		EncodeDuration: MSecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
	conf.EncoderConfig = encoderConf
	logger, _ = conf.Build(zap.AddCallerSkip(1))
}

func getTraceID(ctx context.Context) string {
	var traceID string
	if val, ok := ctx.Value(string(TraceIDKey)).(string); ok {
		traceID = val
	}
	return traceID
}

func WithTraceID(ctx context.Context, traceID string) context.Context {
	return context.WithValue(ctx, TraceIDKey, traceID)
}

func traceID(ctx context.Context, fs []zap.Field) []zap.Field {
	if id := getTraceID(ctx); len(id) != 0 {
		fs = append(fs, zap.String(string(TraceIDKey), id))
	}
	return fs
}

func Debug(ctx context.Context, msg string, fs ...zap.Field) {
	logger.Debug(msg, traceID(ctx, fs)...)
}

func Info(ctx context.Context, msg string, fs ...zap.Field) {
	logger.Info(msg, traceID(ctx, fs)...)
}

func Warning(ctx context.Context, msg string, fs ...zap.Field) {
	logger.Warn(msg, traceID(ctx, fs)...)
}

func Error(ctx context.Context, msg string, fs ...zap.Field) {
	logger.Error(msg, traceID(ctx, fs)...)
}

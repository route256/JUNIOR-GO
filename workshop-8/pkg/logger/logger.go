package logger

import (
	"context"

	"go.uber.org/zap"
)

var defaultLogger *zap.Logger

type ctxKey struct{}

func SetGlobal(logger *zap.Logger) {
	defaultLogger = logger
}

func FromContext(ctx context.Context) *zap.Logger {
	if logger, ok := ctx.Value(ctxKey{}).(*zap.Logger); ok {
		return logger
	}
	return defaultLogger
}

func ToContext(ctx context.Context, logger *zap.Logger) context.Context {
	return context.WithValue(ctx, ctxKey{}, logger)
}

func Infof(ctx context.Context, format string, args ...any) {
	FromContext(ctx).Sugar().Infof(format, args...)
}

func Errorf(ctx context.Context, format string, args ...any) {
	FromContext(ctx).Sugar().Errorf(format, args...)
}

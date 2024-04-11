package logger

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"time"
)

type ctxTimestampKey struct{}

var defaultLogger *slog.Logger

func ContextWithTimestamp(ctx context.Context, ts time.Time) context.Context {
	return context.WithValue(ctx, ctxTimestampKey{}, ts)
}

func TimestampFromContext(ctx context.Context) (time.Time, bool) {
	ts, ok := ctx.Value(ctxTimestampKey{}).(time.Time)
	if !ok {
		return ts, false
	}
	return ts, true
}

type handler struct {
	slog.Handler
}

func (h *handler) Handle(ctx context.Context, r slog.Record) error {
	ts, ok := TimestampFromContext(ctx)
	if ok {
		r.AddAttrs(slog.Any("d", time.Since(ts).String()))
	}

	return h.Handler.Handle(ctx, r)
}

func init() {
	h := handler{
		Handler: slog.NewJSONHandler(os.Stderr, nil),
	}
	defaultLogger = slog.New(&h)
}

func Default() *slog.Logger {
	return defaultLogger
}

func Info(ctx context.Context, msg string, args ...any) {
	Default().Log(ctx, slog.LevelInfo, msg, args...)
}

func Infof(ctx context.Context, format string, args ...any) {
	Default().Log(ctx, slog.LevelInfo, fmt.Sprintf(format, args...))
}

func Error(ctx context.Context, msg string, args ...any) {
	Default().Log(ctx, slog.LevelError, msg, args...)
}

func Errorf(ctx context.Context, format string, args ...any) {
	Default().Log(ctx, slog.LevelError, fmt.Sprintf(format, args...))
}

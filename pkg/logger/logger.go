package logger

import (
	"context"
	"log/slog"
	"os"
)

// Logger is a struct which encapsulates a slog logger.
type Logger struct {
	logger *slog.Logger
}

// NewDebugJsonSlogLogger creates a new logger with debug level and JSON format.
func NewDebugJsonSlogLogger() Logger {
	return Logger{
		logger: slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelDebug.Level(),
		})),
	}
}

// NewInfoJsonSlogLogger creates a new logger with info level and JSON format.
func NewInfoJsonSlogLogger() Logger {
	return Logger{
		logger: slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelInfo.Level(),
		})),
	}
}

// Debug logs a DEBUG level message with context and optional arguments.
func (l Logger) Debug(ctx context.Context, msg string, args ...any) {
	l.logger.DebugContext(ctx, msg, args...)
}

// Info logs an INFO level message with context and optional arguments.
func (l Logger) Info(ctx context.Context, msg string, args ...any) {
	l.logger.InfoContext(ctx, msg, args...)
}

// Warn logs a WARN level message with context and optional arguments.
func (l Logger) Warn(ctx context.Context, msg string, args ...any) {
	l.logger.WarnContext(ctx, msg, args...)
}

// Error logs an ERROR level message with context and optional arguments.
func (l Logger) Error(ctx context.Context, msg string, args ...any) {
	l.logger.ErrorContext(ctx, msg, args...)
}

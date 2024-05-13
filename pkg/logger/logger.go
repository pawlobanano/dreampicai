package logger

import (
	"log/slog"
	"os"
)

var programLevel = new(slog.LevelVar)

// SlogLogger is a logger that uses slog.
type SlogLogger struct{}

// Debug logs a debug message with optional arguments.
func (l *SlogLogger) Debug(msg string, args ...interface{}) {
	slog.Debug(msg, args...)
}

// Info logs an info message with optional arguments.
func (l *SlogLogger) Info(msg string, args ...interface{}) {
	slog.Info(msg, args...)
}

// Warn logs a warning message with optional arguments.
func (l *SlogLogger) Warn(msg string, args ...interface{}) {
	slog.Warn(msg, args...)
}

// Error logs an error message with optional arguments.
func (l *SlogLogger) Error(msg string, args ...interface{}) {
	slog.Error(msg, args...)
}

// InitSlogLogger creates a new slog logger.
func InitSlogLogger() *slog.Logger {
	programLevel.Set(slog.LevelDebug)
	return slog.New(slog.NewJSONHandler(os.Stdout, nil))
}

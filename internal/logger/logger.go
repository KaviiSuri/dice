package logger

import (
	"context"
	"io"
	"log/slog"
	"os"
	"sync/atomic"

	"github.com/dicedb/dice/config"
	"github.com/rs/zerolog"
)

func getLogLevel() slog.Leveler {
	var level slog.Leveler
	switch config.DiceConfig.Server.LogLevel {
	case "debug":
		level = slog.LevelDebug
	case "warn":
		level = slog.LevelWarn
	case "error":
		level = slog.LevelError
	default:
		level = slog.LevelInfo
	}
	return level
}

type Opts struct {
	WithTimestamp bool
}

type Logger struct {
	*slog.Logger
}

var defaultLogger atomic.Pointer[Logger]

func SetDefault(logger *Logger) {
	defaultLogger.Store(logger)
	slog.SetDefault(logger.Logger)
}

func Default() *Logger { return defaultLogger.Load() }

func New(opts Opts) *Logger {
	var writer io.Writer = os.Stderr
	if config.DiceConfig.Server.PrettyPrintLogs {
		writer = zerolog.ConsoleWriter{Out: os.Stderr}
	}
	zerologLogger := zerolog.New(writer)
	if opts.WithTimestamp {
		zerologLogger = zerologLogger.With().Timestamp().Logger()
	}
	logger := &Logger{slog.New(newZerologHandler(zerologLogger))}

	return logger
}

// Fatal logs at [LevelError] level and does os.Exit(1).
func (l *Logger) Fatal(msg string, args ...any) {
	l.FatalContext(context.Background(), msg, args...)
}

// FatalContext logs at [LevelError] level and does os.Exit(1).
func (l *Logger) FatalContext(ctx context.Context, msg string, args ...any) {
	l.ErrorContext(ctx, msg, args...)
	os.Exit(1)
}

// Repeated Functions from slog.Handler

// Debug calls [Logger.Debug] on the default logger.
func Debug(msg string, args ...any) {
	DebugContext(context.Background(), msg, args...)
}

// DebugContext calls [Logger.DebugContext] on the default logger.
func DebugContext(ctx context.Context, msg string, args ...any) {
	Default().DebugContext(ctx, msg, args...)
}

// Info calls [Logger.Info] on the default logger.
func Info(msg string, args ...any) {
	InfoContext(context.Background(), msg, args...)
}

// InfoContext calls [Logger.InfoContext] on the default logger.
func InfoContext(ctx context.Context, msg string, args ...any) {
	Default().InfoContext(ctx, msg, args...)
}

// Warn calls [Logger.Warn] on the default logger.
func Warn(msg string, args ...any) {
	WarnContext(context.Background(), msg, args...)
}

// WarnContext calls [Logger.WarnContext] on the default logger.
func WarnContext(ctx context.Context, msg string, args ...any) {
	Default().WarnContext(ctx, msg, args...)
}

// Error calls [Logger.Error] on the default logger.
func Error(msg string, args ...any) {
	ErrorContext(context.Background(), msg, args...)
}

// ErrorContext calls [Logger.ErrorContext] on the default logger.
func ErrorContext(ctx context.Context, msg string, args ...any) {
	Default().ErrorContext(ctx, msg, args...)
}

// Fatal calls [Logger.Fatal] on the default logger.
func Fatal(msg string, args ...any) {
	FatalContext(context.Background(), msg, args...)
}

// FatalContext calls [Logger.FatalContext] on the default logger.
func FatalContext(ctx context.Context, msg string, args ...any) {
	Default().FatalContext(ctx, msg, args...)
}

// Log calls [Logger.Log] on the default logger.
func Log(ctx context.Context, level slog.Level, msg string, args ...any) {
	Default().Log(ctx, level, msg, args...)
}

// LogAttrs calls [Logger.LogAttrs] on the default logger.
func LogAttrs(ctx context.Context, level slog.Level, msg string, attrs ...slog.Attr) {
	Default().LogAttrs(ctx, level, msg, attrs...)
}

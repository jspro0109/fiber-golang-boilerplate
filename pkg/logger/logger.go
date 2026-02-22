package logger

import (
	"log/slog"
	"os"
	"strings"
)

// Setup configures the default slog logger.
// env "production"/"staging" → JSON output; otherwise → Text output.
// level is one of: debug, info, warn, error (defaults to info).
func Setup(env, level string) {
	lvl := parseLevel(level)
	opts := &slog.HandlerOptions{Level: lvl}

	var handler slog.Handler
	switch strings.ToLower(env) {
	case "production", "staging":
		handler = slog.NewJSONHandler(os.Stdout, opts)
	default:
		handler = slog.NewTextHandler(os.Stdout, opts)
	}

	slog.SetDefault(slog.New(handler))
}

func parseLevel(s string) slog.Level {
	switch strings.ToLower(s) {
	case "debug":
		return slog.LevelDebug
	case "warn", "warning":
		return slog.LevelWarn
	case "error":
		return slog.LevelError
	default:
		return slog.LevelInfo
	}
}

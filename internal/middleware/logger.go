package middleware

import (
	"log/slog"
	"time"

	"github.com/gofiber/fiber/v3"
)

func Logger() fiber.Handler {
	return func(c fiber.Ctx) error {
		start := time.Now()

		err := c.Next()

		latency := time.Since(start)
		status := c.Response().StatusCode()

		var userID int64
		if v := fiber.Locals[int64](c, "user_id"); v != 0 {
			userID = v
		}

		attrs := []slog.Attr{
			slog.String("method", c.Method()),
			slog.String("path", c.Path()),
			slog.Int("status", status),
			slog.Duration("latency", latency),
			slog.String("ip", c.IP()),
			slog.String("request_id", fiber.Locals[string](c, "request_id")),
			slog.Int64("user_id", userID),
			slog.Int("content_length", c.Request().Header.ContentLength()),
			slog.String("user_agent", c.Get("User-Agent")),
			slog.String("query", string(c.Request().URI().QueryString())),
		}

		switch {
		case status >= 500:
			slog.LogAttrs(c.Context(), slog.LevelError, "request", attrs...)
		case status >= 400:
			slog.LogAttrs(c.Context(), slog.LevelWarn, "request", attrs...)
		default:
			slog.LogAttrs(c.Context(), slog.LevelInfo, "request", attrs...)
		}

		return err
	}
}

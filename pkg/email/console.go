package email

import (
	"context"
	"log/slog"
	"strings"
)

type ConsoleSender struct{}

func NewConsoleSender() *ConsoleSender {
	return &ConsoleSender{}
}

func (s *ConsoleSender) Send(_ context.Context, msg Message) error {
	slog.Info("email sent (console driver)",
		slog.String("to", strings.Join(msg.To, ", ")),
		slog.String("subject", msg.Subject),
		slog.String("body", msg.Body),
	)
	return nil
}

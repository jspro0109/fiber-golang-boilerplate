package async

import (
	"log/slog"
	"runtime/debug"
)

// Go runs fn in a new goroutine with panic recovery.
// Any panic is logged and does not crash the process.
func Go(fn func()) {
	go func() {
		defer func() {
			if r := recover(); r != nil {
				slog.Error("async goroutine panicked",
					slog.Any("panic", r),
					slog.String("stack", string(debug.Stack())),
				)
			}
		}()
		fn()
	}()
}

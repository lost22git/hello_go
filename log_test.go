package hello

import (
	"log/slog"
	"testing"
)

func TestLogSlog(t *testing.T) {
	slog.SetLogLoggerLevel(slog.LevelDebug)
	slog.Debug("log a DEBUG msg")
	slog.Info("log a INFO msg")
	slog.Warn("log a WARN msg")
	slog.Error("log a ERROR msg")
}

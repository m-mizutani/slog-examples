package slog_examples

import (
	"errors"
	"os"
	"testing"

	"golang.org/x/exp/slog"
)

func TestSource(t *testing.T) {
	logger := slog.New(slog.HandlerOptions{
		AddSource: true,
	}.NewJSONHandler(os.Stdout))
	logger.Info("hello")
}

func TestLogLevel(t *testing.T) {
	logger := slog.New(slog.HandlerOptions{
		Level: slog.WarnLevel,
	}.NewJSONHandler(os.Stdout))

	logger.Debug("no output")
	logger.Info("no output")
	logger.Warn("WARNING")
	logger.Error("ERROR", errors.New("my error"))
}

func TestReplaceAttr(t *testing.T) {
	logger := slog.New(slog.HandlerOptions{
		ReplaceAttr: func(a slog.Attr) slog.Attr {
			if a.Key == "color" {
				return slog.String("color", "orange")
			}
			return a
		},
	}.NewJSONHandler(os.Stdout))

	logger.Info("hello",
		slog.String("region", "asia"),
		slog.String("color", "blue"),
	)
}

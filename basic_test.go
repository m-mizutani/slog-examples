package slog_examples

import (
	"os"
	"testing"

	"golang.org/x/exp/slog"
)

func TestInfoWithArgs(t *testing.T) {
	slog.Info("hello world!", "name", "blue", "No", 5)
}

func TestJSONHandler(t *testing.T) {
	logger := slog.New(slog.NewJSONHandler(os.Stdout))
	logger.Info("hello structured log", "name", "blue", "No", 5)
}

func TestWith(t *testing.T) {
	logger := slog.New(slog.NewJSONHandler(os.Stdout))
	logger2 := logger.With(
		slog.String("method", "GET"),
		slog.String("path", "/v1/api/hello"),
	)

	logger2.Info("hello")
}

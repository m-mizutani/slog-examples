package slog_examples

import (
	"os"
	"testing"

	"golang.org/x/exp/slog"
)

func TestWithArgs(t *testing.T) {
	logger := slog.New(slog.NewJSONHandler(os.Stdout))
	logger.Info("hello",
		slog.Bool("allowed", true),
		slog.String("name", "blue"),
	)
}

func TestWithMixedArgs(t *testing.T) {
	logger := slog.New(slog.NewJSONHandler(os.Stdout))
	logger.Info("hello",
		slog.Bool("allowed", true),
		slog.String("name", "blue"),
		"version", 3.2,
	)
}

func TestNestedStruct(t *testing.T) {
	data := struct {
		Color  string
		Nested struct {
			Depth int
		}
	}{
		Color: "blue",
		Nested: struct{ Depth int }{
			Depth: 1,
		},
	}

	logger := slog.New(slog.NewJSONHandler(os.Stdout))
	logger.Info("hello", slog.Any("data", data))
}

func TestGroup(t *testing.T) {
	logger := slog.New(slog.NewJSONHandler(os.Stdout))
	logger.Info("hello",
		slog.Group("group1",
			slog.String("color", "blue"),
		),
		slog.Group("group2",
			slog.String("version", "v3.2.0"),
		),
	)
}

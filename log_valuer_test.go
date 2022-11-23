package slog_examples

import (
	"os"
	"strings"
	"testing"

	"golang.org/x/exp/slog"
)

type upperCase string

func (x upperCase) LogValue() slog.Value {
	return slog.StringValue(strings.ToUpper(string(x)))
}

func TestLogValuer(t *testing.T) {
	var s upperCase = "is it small?"
	logger := slog.New(slog.NewJSONHandler(os.Stdout))
	logger.Info("hello", "s", s)
}

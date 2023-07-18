package main

import (
	"errors"
	"os"

	"github.com/zailic/slogr"
	"golang.org/x/exp/slog"
)

// A token is a secret value that grants permissions.
type Token string

// LogValue implements slog.LogValuer.
// It avoids revealing the token.
func (Token) LogValue() slog.Value {
	return slog.StringValue("REDACTED_TOKEN")
}

// This example demonstrates a Value that replaces itself
// with an alternative representation to avoid revealing secrets.
func main() {
	t := Token("shhhh!")
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{}))
	var log slogr.Logger = slogr.New(logger)
	err := errors.New("permission denied")
	log.Info("permission granted", "user", "Perry", "token", t)
	log.Error(err, "error occurred")
	log2 := log.WithName("subsystem").WithValues("type", "authentication")
	log2.Info("permission granted", "user", "Perry", "token", t)
}

// Pacakge: slogr provides an implementation of logr(github.com/go-logr/logr)
// interfaces backed by slog("golang.org/x/exp/slog).
//
// Usage:
// A logr.Logger instance can be created from slog.Logger by slogr.New method:
//   	slogLogger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
//  	logger := slogr.New(slodLogger)

package main

import (
	"context"

	logr "github.com/go-logr/logr"
	"golang.org/x/exp/slog"
)

// Logger is an alias of logr.Logger.
type Logger = logr.Logger

var (
	_ logr.LogSink          = &LogSink{}
	_ logr.CallDepthLogSink = &LogSink{}
)

// LogSink is a logr.LogSink backed by slog.
type LogSink struct {
	logger *slog.Logger
}

// WithCallDepth returns the same logr.LogSink instance because slog.Logger
// doesn't support a configurable call depth.
func (s *LogSink) WithCallDepth(depth int) logr.LogSink {
	// underlying slog.Logger doesn't support a configuragle call depth
	// despite it's logging methods contains a logic that calls runtime.Callers
	// but the "skip" argument is hardcoded to 3.
	return s
}

// Enabled reports whether the backend logger is enabled at the specified level.
// The level value should be a compatible slog level value.(e.g. slog.LevelDebug)
func (s *LogSink) Enabled(level int) bool {
	return s.logger.Enabled(context.Background(), slog.Level(level))
}

// Error logs an error, with the given message and key/value pairs as context.
func (s *LogSink) Error(err error, msg string, keysAndValues ...interface{}) {
	keysAndValues = append(keysAndValues, "error", err)
	s.logger.Error(msg, keysAndValues...)
}

// Info implements logr.LogSink.
func (s *LogSink) Info(level int, msg string, keysAndValues ...interface{}) {
	s.logger.Log(context.Background(), slog.Level(level), msg, keysAndValues...)
}

// Init doesn't do anything because slog.Logger doesn't expose a way to set
// the logr.RuntimeInfo.CallDepth.
func (s *LogSink) Init(info logr.RuntimeInfo) {
}

// WithName returns a new logr.LogSink that starts a group having the given name.
// The keys of all attributes added to the new logr.LogSink will be qualified by
// the given name. (How that qualification happens depends on the
// [Handler.WithGroup] method of the logr.LogSink's Handler.)
func (s LogSink) WithName(name string) logr.LogSink {
	s.logger = s.logger.WithGroup(name)
	return &s
}

// WithValues returns a new logr.LogSink with additional key/value pairs set.
func (s LogSink) WithValues(keysAndValues ...interface{}) logr.LogSink {
	s.logger = s.logger.With(keysAndValues...)
	return &s
}

// NewLogSink returns a new logr.LogSink backed by slog.
func NewLogSink(logger *slog.Logger) logr.LogSink {
	return &LogSink{
		logger: logger,
	}
}

// New returns a new logr.Logger backed by slog.
func New(logger *slog.Logger) Logger {
	return logr.New(NewLogSink(logger))
}

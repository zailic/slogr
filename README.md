# Slogr
Slogr provides an implementation of [logr](github.com/go-logr/logr) interfaces backed by [slog]("golang.org/x/exp/slog).

## Usage example
```go
slogLogger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{}))
var log logr.Logger = slogr.New(slogLogger)
log.Info("Access granted.", "user", "john.doe" )
```

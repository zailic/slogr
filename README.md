# Slogr
Slogr provides an implementation of [logr](https://github.com/go-logr/logr) interfaces backed by [slog](https://golang.org/x/exp/slog).

## Usage example
```go
slogLogger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{}))
var log logr.Logger = slogr.New(slogLogger)
log.Info("Access granted.", "user", "john.doe" )
```

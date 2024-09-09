package logger

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5/middleware"
	"go.uber.org/zap"
)

// RequestLogger returns a logger handler..
func RequestLogger() func(next http.Handler) http.Handler {
	return middleware.RequestLogger(&LogFormatter{})
}

// LogEntry records the final log when a request completes.
type LogEntry struct {
	logger *zap.SugaredLogger
}

func (l *LogEntry) Write(status, bytes int, header http.Header, elapsed time.Duration, extra interface{}) {
	l.logger.With(
		"status", status,
		"resp_bytes_length", bytes,
		"elapsed", fmt.Sprint(elapsed),
	).Info()
}

func (l *LogEntry) Panic(v interface{}, stack []byte) {
	l.logger = l.logger.With(
		"stack", string(stack),
		"panic", fmt.Sprintf("%+v", v),
	)
}

// LogFormatter initiates the beginning of a new LogEntry per request.
type LogFormatter struct {
}

func (l *LogFormatter) NewLogEntry(r *http.Request) middleware.LogEntry {
	scheme := "http"
	if r.TLS != nil {
		scheme = "https"
	}

	log := FromContext(r.Context()).With(
		"http_scheme", scheme,
		"uri", fmt.Sprintf("%s://%s%s", scheme, r.Host, r.RequestURI),
		"http_proto", r.Proto,
		"http_method", r.Method,
		"remote_addr", r.RemoteAddr,
		"user_agent", r.UserAgent(),
	)

	return &LogEntry{logger: log}
}

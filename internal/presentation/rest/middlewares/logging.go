package middlewares

import (
	"log/slog"
	"net/http"
	"time"
)

func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now().UTC()

		log := slog.With(
			slog.String("method", r.Method),
			slog.String("path", r.URL.Path),
		)

		log.InfoContext(r.Context(), "request started")

		next.ServeHTTP(w, r)

		log.InfoContext(r.Context(), "request finished", slog.Duration("latency", time.Since(start)))
	})
}

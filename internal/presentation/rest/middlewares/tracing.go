package middlewares

import (
	"context"
	"net/http"

	"github.com/google/uuid"
)

// TODO: implement trace id from env
func Tracing(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestID := r.Header.Get("X-Trace-ID")
		if requestID == "" {
			requestID = uuid.NewString()
		}

		ctx := context.WithValue(r.Context(), "trace_id", requestID)

		w.Header().Set("X-Trace-ID", requestID)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

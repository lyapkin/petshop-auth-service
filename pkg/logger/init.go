package logger

import (
	"context"
	"log/slog"
	"os"

	"github.com/lyapkin/shop/auth/config"
)

type contextHandler struct {
	slog.Handler
}

func New(env config.Env) *slog.Logger {
	// TODO: seperate out by env
	baseHandler := slog.NewTextHandler(os.Stdout, nil)
	logger := slog.New(&contextHandler{baseHandler})
	return logger
}

func (h *contextHandler) Handle(ctx context.Context, r slog.Record) error {
	// TODO: trace_id from env
	if traceID, ok := ctx.Value("trace_id").(string); ok {
		r.AddAttrs(slog.String("trace_id", traceID))
	}
	return h.Handler.Handle(ctx, r)
}

func (h *contextHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return &contextHandler{h.Handler.WithAttrs(attrs)}
}

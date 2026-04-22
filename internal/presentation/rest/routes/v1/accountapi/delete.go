package accountapi

import (
	"log/slog"
	"net/http"

	"github.com/google/uuid"
	"github.com/lyapkin/shop/auth/internal/app/domain"
	"github.com/lyapkin/shop/auth/internal/presentation/rest/utils/response"
)

func (h *handler) delete(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(r.PathValue("id"))
	if err != nil {
		slog.WarnContext(r.Context(), "invalid url parameter", slog.String("err", err.Error()))
		response.ResWithError(w, &domain.AppError{
			Code:     domain.ErrNotValid,
			Message:  "invalid url parameter",
			Internal: err,
		})
		return
	}

	if err := h.uc.Delete(r.Context(), id); err != nil {
		response.ResWithError(w, err)
		return
	}

	response.ResWithSuccess(w, http.StatusNoContent, nil)
}

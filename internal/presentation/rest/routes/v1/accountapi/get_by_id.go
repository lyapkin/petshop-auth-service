package accountapi

import (
	"log/slog"
	"net/http"

	"github.com/google/uuid"
	"github.com/lyapkin/shop/auth/internal/app/domain"
	"github.com/lyapkin/shop/auth/internal/presentation/rest/utils/response"
)

func (h *handler) getByID(w http.ResponseWriter, r *http.Request) {
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

	result, err := h.uc.GetByID(r.Context(), id)
	if err != nil {
		response.ResWithError(w, err)
		return
	}

	response.ResWithSuccess(w, http.StatusOK, result)
}

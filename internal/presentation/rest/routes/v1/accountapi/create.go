package accountapi

import (
	"log/slog"
	"net/http"

	"github.com/lyapkin/shop/auth/internal/app/domain"
	"github.com/lyapkin/shop/auth/internal/presentation/rest/utils/request"
	"github.com/lyapkin/shop/auth/internal/presentation/rest/utils/response"
)

func (h *handler) create(w http.ResponseWriter, r *http.Request) {
	var input domain.Account
	if err := request.ParseBody(r.Body, &input); err != nil {
		slog.WarnContext(r.Context(), "failed to parse request body", slog.String("err", err.Error()))
		response.ResWithError(w, &domain.AppError{
			Code:     domain.ErrNotValid,
			Message:  "invalid request body",
			Internal: err,
		})
		return
	}

	result, err := h.uc.Create(r.Context(), &input)
	if err != nil {
		response.ResWithError(w, err)
		return
	}

	response.ResWithSuccess(w, http.StatusCreated, result)
}

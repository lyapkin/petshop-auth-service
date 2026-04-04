package authapi

import (
	"log/slog"
	"net/http"

	"github.com/lyapkin/shop/auth/internal/app/domain"
	"github.com/lyapkin/shop/auth/internal/presentation/rest/utils/request"
	"github.com/lyapkin/shop/auth/internal/presentation/rest/utils/response"
)

func (h *handler) refreshToken(w http.ResponseWriter, r *http.Request) {
	var input domain.RefreshToken
	if err := request.ParseBody(r.Body, &input); err != nil {
		slog.WarnContext(r.Context(), "failed to parse refresh body", slog.String("err", err.Error()))
		response.ResWithError(w, &domain.AppError{
			Code:     domain.ErrNotValid,
			Message:  "invalid request body",
			Internal: err,
		})
		return
	}

	token, err := h.uc.Refresh(r.Context(), input.Token)
	if err != nil {
		response.ResWithError(w, err)
		return
	}

	response.ResWithSuccess(w, http.StatusOK, token)
}

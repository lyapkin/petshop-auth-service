package authapi

import (
	"log/slog"
	"net/http"

	"github.com/lyapkin/shop/auth/internal/app/domain"
	"github.com/lyapkin/shop/auth/internal/app/dto"
	"github.com/lyapkin/shop/auth/internal/presentation/rest/utils/request"
	"github.com/lyapkin/shop/auth/internal/presentation/rest/utils/response"
)

func (h *handler) register(w http.ResponseWriter, r *http.Request) {
	var input dto.RegistrationInput
	if err := request.ParseBody(r.Body, &input); err != nil {
		slog.WarnContext(r.Context(), "registration input parsing failed", slog.String("err", err.Error()))
		response.ResWithError(w, &domain.AppError{
			Code:     domain.ErrNotValid,
			Message:  "invalid request body",
			Internal: err,
		})
		return
	}

	account := domain.Account{
		Email:    input.Email,
		Password: input.Password,
	}

	err := h.uc.Register(r.Context(), &account)
	if err != nil {
		response.ResWithError(w, err)
		return
	}

	response.ResWithSuccess(w, http.StatusOK, nil)
}

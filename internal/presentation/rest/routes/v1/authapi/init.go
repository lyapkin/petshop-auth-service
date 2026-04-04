package authapi

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/lyapkin/shop/auth/internal/app/usecases/auth"
)

type handler struct {
	uc *auth.Usecase
}

func New(uc *auth.Usecase) http.Handler {
	handler := &handler{
		uc: uc,
	}

	r := chi.NewRouter()

	r.Post("/register", handler.register)
	r.Post("/login", handler.login)
	r.Post("/logout", handler.logout)
	r.Post("/refresh", handler.refreshToken)

	return r
}

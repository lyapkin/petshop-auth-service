package rest

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/lyapkin/shop/auth/internal/app/usecases/auth"
	v1 "github.com/lyapkin/shop/auth/internal/presentation/rest/routes/v1"
)

func New(auth *auth.Usecase) http.Handler {
	r := chi.NewRouter()

	r.Route("/api", func(r chi.Router) {
		r.Mount("/v1", v1.New(auth))
	})

	return r
}

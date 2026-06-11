package accountapi

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/lyapkin/shop/auth/internal/app/interfaces"
	"github.com/lyapkin/shop/auth/internal/app/usecases/account"
	"github.com/lyapkin/shop/auth/internal/presentation/rest/middlewares"
)

type handler struct {
	uc             *account.Usecase
	tokenValidator interfaces.TokenValidator
}

func New(uc *account.Usecase, tokenValidator interfaces.TokenValidator) http.Handler {
	handler := &handler{
		uc: uc,
	}

	r := chi.NewRouter()

	auth := middlewares.Auth(tokenValidator)

	r.With(auth).Get("/", handler.list)
	r.Post("/", handler.create)
	r.Get("/{id}", handler.getByID)
	r.Put("/{id}", handler.update)
	r.Delete("/{id}", handler.delete)

	return r
}

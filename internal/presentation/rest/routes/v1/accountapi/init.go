package accountapi

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/lyapkin/shop/auth/internal/app/usecases/account"
)

type handler struct {
	uc *account.Usecase
}

func New(uc *account.Usecase) http.Handler {
	handler := &handler{
		uc: uc,
	}

	r := chi.NewRouter()

	r.Get("/", handler.list)
	r.Post("/", handler.create)
	r.Get("/{id}", handler.getByID)
	r.Put("/{id}", handler.update)
	r.Delete("/{id}", handler.delete)

	return r
}

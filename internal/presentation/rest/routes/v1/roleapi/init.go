package roleapi

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/lyapkin/shop/auth/internal/app/usecases/role"
)

type handler struct {
	uc *role.Usecase
}

func New(uc *role.Usecase) http.Handler {
	handler := &handler{
		uc: uc,
	}

	r := chi.NewRouter()

	r.Post("/", handler.create)
	r.Get("/", handler.list)
	r.Put("/{id}", handler.update)
	r.Delete("/{id}", handler.delete)

	return r
}

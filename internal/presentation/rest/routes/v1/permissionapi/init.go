package permissionapi

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/lyapkin/shop/auth/internal/app/usecases/permission"
)

type handler struct {
	uc *permission.Usecase
}

func New(uc *permission.Usecase) http.Handler {
	handler := handler{
		uc: uc,
	}

	r := chi.NewRouter()

	r.Get("/", handler.list)
	r.Post("/", handler.create)
	r.Delete("/{id}", handler.delete)
	r.Put("/{id}", handler.update)

	return r
}

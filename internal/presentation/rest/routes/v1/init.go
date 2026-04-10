package v1

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/lyapkin/shop/auth/internal/app/usecases/auth"
	"github.com/lyapkin/shop/auth/internal/app/usecases/role"
	"github.com/lyapkin/shop/auth/internal/presentation/rest/routes/v1/authapi"
	"github.com/lyapkin/shop/auth/internal/presentation/rest/routes/v1/roleapi"
)

func New(auth *auth.Usecase, role *role.Usecase) http.Handler {
	r := chi.NewRouter()

	r.Mount("/auth", authapi.New(auth))
	r.Mount("/roles", roleapi.New(role))

	return r
}

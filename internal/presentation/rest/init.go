package rest

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/lyapkin/shop/auth/internal/app/interfaces"
	"github.com/lyapkin/shop/auth/internal/app/usecases/account"
	"github.com/lyapkin/shop/auth/internal/app/usecases/auth"
	"github.com/lyapkin/shop/auth/internal/app/usecases/permission"
	"github.com/lyapkin/shop/auth/internal/app/usecases/role"
	mw "github.com/lyapkin/shop/auth/internal/presentation/rest/middlewares"
	v1 "github.com/lyapkin/shop/auth/internal/presentation/rest/routes/v1"
)

func New(
	auth *auth.Usecase,
	role *role.Usecase,
	permission *permission.Usecase,
	account *account.Usecase,
	tokenValidator interfaces.TokenValidator,
) http.Handler {
	r := chi.NewRouter()

	r.Use(mw.Tracing)
	r.Use(mw.Logging)

	r.Route("/api", func(r chi.Router) {
		r.Mount("/v1", v1.New(auth, role, permission, account, tokenValidator))
	})

	return r
}

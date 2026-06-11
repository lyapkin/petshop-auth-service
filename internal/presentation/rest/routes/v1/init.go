package v1

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/lyapkin/shop/auth/internal/app/interfaces"
	"github.com/lyapkin/shop/auth/internal/app/usecases/account"
	"github.com/lyapkin/shop/auth/internal/app/usecases/auth"
	"github.com/lyapkin/shop/auth/internal/app/usecases/permission"
	"github.com/lyapkin/shop/auth/internal/app/usecases/role"
	"github.com/lyapkin/shop/auth/internal/presentation/rest/routes/v1/accountapi"
	"github.com/lyapkin/shop/auth/internal/presentation/rest/routes/v1/authapi"
	"github.com/lyapkin/shop/auth/internal/presentation/rest/routes/v1/permissionapi"
	"github.com/lyapkin/shop/auth/internal/presentation/rest/routes/v1/roleapi"
)

func New(
	auth *auth.Usecase,
	role *role.Usecase,
	permission *permission.Usecase,
	account *account.Usecase,
	tokenValidator interfaces.TokenValidator,
) http.Handler {
	r := chi.NewRouter()

	r.Mount("/auth", authapi.New(auth))
	r.Mount("/roles", roleapi.New(role))
	r.Mount("/permissions", permissionapi.New(permission))
	r.Mount("/accounts", accountapi.New(account, tokenValidator))

	r.Mount("/health", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	}))

	return r
}

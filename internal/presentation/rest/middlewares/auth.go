package middlewares

import (
	"context"
	"errors"
	"net/http"
	"strings"

	"github.com/lyapkin/shop/auth/internal/app/domain"
	"github.com/lyapkin/shop/auth/internal/app/interfaces"
	"github.com/lyapkin/shop/auth/internal/presentation/rest/utils/response"
)

func Auth(validator interfaces.TokenValidator) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authHeaders := r.Header.Get("Authorization")
			if len(authHeaders) == 0 {
				response.ResWithError(w, &domain.AppError{
					Code:     domain.ErrNotAuthenticated,
					Message:  "not authenticated",
					Internal: errors.New("no token"),
				})
				return
			}

			authParts := strings.Split(authHeaders, " ")
			if len(authParts) != 2 || strings.ToLower(authParts[0]) != "bearer" {
				response.ResWithError(w, &domain.AppError{
					Code:     domain.ErrNotAuthenticated,
					Message:  "invalid authorization format. expected 'Bearer <Token>'",
					Internal: errors.New("invalid format of authorization header"),
				})
				return
			}

			token := authParts[1]

			claims, err := validator.ValidateToken(token)
			if err != nil {
				response.ResWithError(w, err)
				return
			}

			ctx := context.WithValue(r.Context(), "user", claims)

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

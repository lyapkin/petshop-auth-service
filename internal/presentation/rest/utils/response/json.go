package response

import (
	"encoding/json"
	"net/http"

	"github.com/lyapkin/shop/auth/internal/app/domain"
)

func resWithJSON(w http.ResponseWriter, code int, payload any) error {
	w.Header().Add("Content-Type", "application/json")

	data, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		res := responseError{
			Code:    string(domain.ErrInternal),
			Message: domain.InternalErrorMessage,
		}
		w.Write(res.toFallbackJSON())
		return err
	}

	w.WriteHeader(code)
	w.Write(data)

	return nil
}

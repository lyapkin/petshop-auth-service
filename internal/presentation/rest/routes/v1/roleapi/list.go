package roleapi

import (
	"net/http"

	"github.com/lyapkin/shop/auth/internal/presentation/rest/utils/response"
)

func (h *handler) list(w http.ResponseWriter, r *http.Request) {
	result, err := h.uc.List(r.Context())
	if err != nil {
		response.ResWithError(w, err)
		return
	}

	response.ResWithSuccess(w, http.StatusOK, result)
}

package response

import "net/http"

func ResWithSuccess(w http.ResponseWriter, code int, payload any) {
	res := responseSuccess{
		Data: payload,
	}

	resWithJSON(w, code, res)
}

type responseSuccess struct {
	Data any `json:"data"`
}

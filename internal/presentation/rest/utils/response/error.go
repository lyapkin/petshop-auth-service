package response

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/lyapkin/shop/auth/internal/app/domain"
)

var httpErrCode = map[domain.ErrorCode]int{
	domain.ErrNotFound:         http.StatusNotFound,
	domain.ErrDuplicate:        http.StatusBadRequest,
	domain.ErrNotValid:         http.StatusBadRequest,
	domain.ErrIDExists:         http.StatusInternalServerError,
	domain.ErrInternal:         http.StatusInternalServerError,
	domain.ErrNotAuthenticated: http.StatusUnauthorized,
}

func ResWithError(w http.ResponseWriter, err error) {
	errCode := domain.ErrInternal
	message := domain.InternalErrorMessage

	var appErr *domain.AppError
	if errors.As(err, &appErr) {
		errCode = appErr.Code
		message = appErr.Message
	}

	res := responseError{
		Code:    string(errCode),
		Message: message,
	}

	resWithJSON(w, httpErrCode[errCode], res)
}

type responseError struct {
	Code    string `json:"code"`
	Message string `json:"msg"`
}

func (e responseError) toFallbackJSON() []byte {
	return []byte(`{` +
		`"code":` + strconv.Quote(e.Code) + `,` +
		`"message":` + strconv.Quote(e.Message) + `,` +
		`}`)
}

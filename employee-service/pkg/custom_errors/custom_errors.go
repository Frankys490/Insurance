package custom_errors

import (
	"github.com/valyala/fasthttp"
)

type ErrHttp struct {
	Code    int
	Message string
}

var (
	ErrEmptyFields        = &ErrHttp{Code: fasthttp.StatusUnprocessableEntity, Message: "fields can't be empty"}
	ErrWrongUserInputData = &ErrHttp{Code: fasthttp.StatusUnprocessableEntity, Message: "user with this username doesn't exist"}
	ErrWrongFormatData    = &ErrHttp{Code: fasthttp.StatusUnprocessableEntity, Message: "Wrong format of input data"}
	ErrInternal           = &ErrHttp{Code: fasthttp.StatusInternalServerError, Message: "Something went wrong"}
	ErrWrongMethod        = &ErrHttp{Code: fasthttp.StatusMethodNotAllowed, Message: "wrong method"}
	ErrNotFound           = &ErrHttp{Code: fasthttp.StatusNotFound, Message: "no records in DB"}
	ErrUserExist          = &ErrHttp{Code: fasthttp.StatusUnprocessableEntity, Message: "user already exist"}
	ErrUnauthorized       = &ErrHttp{Code: fasthttp.StatusUnauthorized, Message: "unauthorized"}
)

func (e *ErrHttp) Error() string {
	if e == nil {
		return ""
	}

	return e.Message
}

func New(code int, message string) *ErrHttp {
	return &ErrHttp{Code: code, Message: message}
}

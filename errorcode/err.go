package errorcode

import "errors"

type (
	ErrType struct {
		Code int
		Err  error
	}

	BusinessErrorWrapper struct {
		Code  int    `json:"code"`
		Error string `json:"error"`
	}
)

func (e ErrType) Error() string {
	return e.Err.Error()
}

func (e ErrType) GetCode() int {
	return e.Code
}

var (
	TokenExpired = ErrType{9000, errors.New("token is expired")}
	RequestErr   = ErrType{9001, errors.New("token is expired")}
)

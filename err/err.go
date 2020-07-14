package err

import "errors"

var (
	ErrInvalidRequestType = errors.New("invalid request parameter")
)

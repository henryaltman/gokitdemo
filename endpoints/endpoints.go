package endpoints

import (
	"context"
	"errors"
	"gokitdemo/services"
	"strings"

	"github.com/go-kit/kit/endpoint"
)

var (
	ErrInvalidRequestType = errors.New("invalid request parameter")
)

// ArithmeticRequest define request struct
type BasicRequest struct {
	RequestType string `json:"request_type"`
	A           int    `json:"a"`
	B           int    `json:"b"`
}

// ArithmeticResponse define response struct
type BasicResponse struct {
	Result int   `json:"result"`
	Error  error `json:"error"`
}

// MakeArithmeticEndpoint make endpoint
func MakeBasicEndpoint(svc services.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(BasicRequest)

		var (
			res, a, b int
			calError  error
		)

		a = req.A
		b = req.B

		if strings.EqualFold(req.RequestType, "Add") {
			res = svc.Add(a, b)
		} else if strings.EqualFold(req.RequestType, "Substract") {
			res = svc.Subtract(a, b)
		} else if strings.EqualFold(req.RequestType, "Multiply") {
			res = svc.Multiply(a, b)
		} else if strings.EqualFold(req.RequestType, "Divide") {
			res, calError = svc.Divide(a, b)
		} else {
			return nil, ErrInvalidRequestType
		}

		return BasicResponse{Result: res, Error: calError}, nil
	}
}

package endpoints

import (
	"context"
	"encoding/json"
	"errors"
	"gokitdemo/dto"
	"gokitdemo/services"

	"github.com/go-kit/kit/endpoint"
)

const (
	RequestTypeAdd      = "add"
	RequestTypeSub      = "sub"
	RequestTypeMultiply = "multiply"
	RequestTypeDivide   = "divide"
)

// MakeBasicEndpoint make endpoint
func MakeBasicEndpoint(svc services.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(dto.BasicRequest)

		dataByte, _ := json.Marshal(req.Data)
		if req.RequestType == RequestTypeAdd {
			addReq := dto.AddRequest{}
			json.Unmarshal(dataByte, &addReq)
			response, err = svc.Add(ctx, addReq)
		} else {
			response, err = nil, errors.New("request router method not found")
		}

		result := dto.BasicResponse{}
		result.Data = response
		if err != nil {
			result.Msg = err.Error()
		}
		return result, nil
	}
}

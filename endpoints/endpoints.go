package endpoints

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"gokitdemo/core"
	"gokitdemo/dto"
	"gokitdemo/services"

	"github.com/go-kit/kit/endpoint"
)

// MakeBasicEndpoint make endpoint
func MakeBasicEndpoint(svc services.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(dto.BasicRequest)
		fmt.Println("req", req)
		result := dto.BasicResponse{}
		if callResult := core.CallReflect(svc, req.Path, ctx, req.Data); callResult != nil {
			callResultByte, _ := json.Marshal(callResult[0].Interface())
			br := services.BaseResponse{}
			err = json.Unmarshal(callResultByte, &br)
			if err != nil {
				result.Msg = err.Error()
			}
			if br.Err != nil {
				result.Msg = br.Err.Error()
			}
			result.Data = br.Rs
		} else {
			response, err = nil, errors.New(fmt.Sprintf("not found method %s", req.Path))
		}
		if err != nil && result.Msg == "" {
			result.Msg = err.Error()
		}
		response = result
		return
	}
}

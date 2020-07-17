package endpoints

import (
	"context"
	"errors"
	"fmt"
	"gokitdemo/auth"
	"gokitdemo/core"
	"gokitdemo/dto"
	"gokitdemo/services"

	"github.com/go-kit/kit/endpoint"
)

// MakeBasicEndpoint make endpoint
func MakeBasicEndpoint(svc services.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {

		result := dto.BasicResponse{}
		//verify token
		_, err = auth.VerifyToken(ctx)
		if err != nil {
			result.Msg = err.Error()
			fmt.Println(fmt.Sprintf("MakeBasicEndpoint err %v", result))
			response = result
			return response, nil
		}
		//verify token

		if request == nil {
			result.Msg = "request error"
			response = result
			return
		}
		req := request.(dto.BasicRequest)
		if callResult := core.CallReflect(svc, req.Path, ctx, req.Request); callResult != nil {
			br := callResult[0].Interface().(services.BaseResponse)
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

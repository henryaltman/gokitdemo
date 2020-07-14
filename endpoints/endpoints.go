package endpoints

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"gokitdemo/dto"
	"gokitdemo/services"
	"reflect"

	"github.com/go-kit/kit/endpoint"
)

const (
	RequestTypeAdd      = "Add"
	RequestTypeSub      = "sub"
	RequestTypeMultiply = "multiply"
	RequestTypeDivide   = "divide"
)

// MakeBasicEndpoint make endpoint
//todo
func MakeBasicEndpoint(svc services.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(dto.BasicRequest)
		fmt.Println("req", req)
		result := dto.BasicResponse{}
		if callResult := callReflect(svc, req.RequestType, ctx, req.Data); callResult != nil {
			result.Data = callResult[0]
			//result.Data = 1
		} else {
			response, err = nil, errors.New(fmt.Sprintf("not found method %s", req.RequestType))
		}
		if err != nil && result.Msg == "" {
			result.Msg = err.Error()
		}
		response = result
		return
	}
}

func callReflect(any interface{}, name string, args ...interface{}) []reflect.Value {
	inputs := make([]reflect.Value, len(args))
	for i, _ := range args {
		inputs[i] = reflect.ValueOf(args[i])
	}
	fmt.Println("inputs", inputs)

	if v := reflect.ValueOf(any).MethodByName(name); v.String() == "<invalid Value>" {
		fmt.Println("<invalid Value>")
		return nil
	} else {
		ret := v.Call(inputs)
		retByte, _ := json.Marshal(ret)
		fmt.Println("ret", string(retByte))
		return ret

	}
}

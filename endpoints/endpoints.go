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
		result := dto.BasicResponse{}
		if callResult := callReflect(svc, req.RequestType, ctx, req.Data); callResult != nil {
			//err = callResult[1].(error)
			callResultByte, _ := json.Marshal(callResult[0])
			fmt.Println("callResultByte", string(callResultByte))
			result.Data = callResult[0]
		} else {
			response, err = nil, errors.New(fmt.Sprintf("not found method %s", req.RequestType))
		}
		//fmt.Println("response", response)

		if err != nil {
			result.Msg = err.Error()
		}
		return result, nil
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
		fmt.Print(ret)
		return ret
	}
}

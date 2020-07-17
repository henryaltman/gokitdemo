package transports

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"gokitdemo/auth"
	"gokitdemo/core"
	"gokitdemo/errorcode"
	"gokitdemo/router"
	"io/ioutil"
	//"log"
	"net/http"

	"gokitdemo/dto"

	//"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/endpoint"
	//kitEndpoint "github.com/go-kit/kit/endpoint"
	"github.com/gorilla/mux"

	"github.com/go-kit/kit/log"

	kitJwt "github.com/go-kit/kit/auth/jwt"
	kitHttp "github.com/go-kit/kit/transport/http"
)

var (
	ErrorBadRequest = errors.New("invalid request parameter")
)

type TransportBaseReqquesst struct {
}

//DecodeDefaultRequest is default service request
func (tbr *TransportBaseReqquesst) DecodeDefaultRequest(data []byte) (dto.BasicRequest, error) {
	return dto.BasicRequest{Path: "Default", RequestId: "xxx", Request: dto.DefaultRequest{}}, nil
}

//DecodeAddRequest is add service request
func (tbr *TransportBaseReqquesst) DecodeAddRequest(data []byte) (dto.BasicRequest, error) {
	request := dto.AddRequest{}
	err := json.Unmarshal(data, &request)
	if err != nil {
		return dto.BasicRequest{}, err
	}
	return dto.BasicRequest{Path: "Add", RequestId: "xxx", Request: request}, nil
}

//DecodeGETRequest is deeal get request
func (tbr *TransportBaseReqquesst) DecodeGETRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	//首字母设为大写
	requestMethodName := fmt.Sprintf("Decode%sRequest", fmt.Sprintf("%v", ctx.Value(auth.HttpPATH)))
	data := []byte{}
	if callResult := core.CallReflect(tbr, requestMethodName, data); callResult != nil {
		return callResult[0].Interface(), nil
	}
	return nil, errors.New("read body data error")
}

//DecodePOSTRequest is deal post request
func (tbr *TransportBaseReqquesst) DecodePOSTRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	bodyBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, errors.New("read body data error")
	}
	routerPath := ctx.Value(auth.HttpPATH)
	requestMethodName := fmt.Sprintf("Decode%sRequest", fmt.Sprintf("%v", routerPath))
	if callResult := core.CallReflect(tbr, requestMethodName, bodyBytes); callResult != nil {
		return callResult[0].Interface(), nil
	}
	return nil, errors.New("read body data error")
}

// DecodeBasicRequest decode request params to struct
func DecodeBasicRequest(ctx context.Context, r *http.Request) (interface{}, error) {

	methodName := fmt.Sprintf("Decode%sRequest", r.Method)
	tbr := &TransportBaseReqquesst{}
	if callResult := core.CallReflect(tbr, methodName, ctx, r); callResult != nil {
		if callResult[1].Interface() != nil {
			return callResult[0].Interface(), nil
		}
		return callResult[0].Interface(), nil
	}
	return nil, ErrorBadRequest
}

// EncodeBasicResponse encode response to return
func EncodeBasicResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Access-Control-Allow-Origin", "*")                                                                          //允许访问所有域
	w.Header().Add("Access-Control-Allow-Headers", "Origin, Content-Type, Authorization, tk, encrypt, authorization, platform") //自定义header头
	w.Header().Set("Access-Control-Allow-Credentials", "true")                                                                  //允许
	w.Header().Add("Access-Control-Allow-Methods", "GET, POST, OPTIONS")                                                        //允许接受
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	if f, ok := response.(endpoint.Failer); ok && f.Failed() != nil {
		businessErrorEncoder(f.Failed(), w)
		return nil
	}

	return json.NewEncoder(w).Encode(response)
}

func encodeError(_ context.Context, err error, w http.ResponseWriter) {

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	customErr, ok := err.(errorcode.ErrType)
	if ok {
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(errorcode.BusinessErrorWrapper{
			Error: customErr.Error(), Code: customErr.GetCode(),
		})
	} else {
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode(map[string]interface{}{"code": 9100, "error": err.Error()})
	}
}

// businessErrorEncoder encode business logic error. for example NotExistErr
func businessErrorEncoder(err error, w http.ResponseWriter) {

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	customErr := err.(errorcode.ErrType)
	_ = json.NewEncoder(w).Encode(errorcode.BusinessErrorWrapper{
		Error: customErr.Error(), Code: customErr.GetCode(),
	})
}

// MakeHttpHandler make http handler use mux
func MakeKitHttpHandler(_ context.Context, endpoint endpoint.Endpoint, logger log.Logger) http.Handler {
	r := mux.NewRouter()
	options := []kitHttp.ServerOption{
		kitHttp.ServerErrorLogger(logger),
		kitHttp.ServerErrorEncoder(encodeError),
		kitHttp.ServerBefore(
			kitJwt.HTTPToContext(), auth.HTTPToContext(),
			auth.LangHTTPToContext(), auth.AuthorizationHTTPToContext()),
	}
	//开始初始化路由
	for _, routerMap := range router.Router {
		r.Methods(routerMap.Method).Path(routerMap.Path).Handler(kitHttp.NewServer(
			endpoint,
			DecodeBasicRequest,
			EncodeBasicResponse,
			options...,
		))
	}
	//结束初始化路由
	return r
}

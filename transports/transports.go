package transports

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"gokitdemo/auth"
	"gokitdemo/core"
	"io/ioutil"
	//"log"
	"net/http"

	"gokitdemo/dto"

	//"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/endpoint"
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
	return json.NewEncoder(w).Encode(response)
}

// MakeHttpHandler make http handler use mux
func MakeKitHttpHandler(_ context.Context, endpoint endpoint.Endpoint, logger log.Logger) http.Handler {
	r := mux.NewRouter()
	options := []kitHttp.ServerOption{
		kitHttp.ServerErrorLogger(logger),
		kitHttp.ServerErrorEncoder(kitHttp.DefaultErrorEncoder),
		kitHttp.ServerBefore(
			kitJwt.HTTPToContext(), auth.HTTPToContext(),
			auth.LangHTTPToContext(), auth.AuthorizationHTTPToContext()),
	}
	r.Methods("GET").Path("/").Handler(kitHttp.NewServer(
		endpoint,
		DecodeBasicRequest,
		EncodeBasicResponse,
		options...,
	))

	r.Methods("POST").Path("/add/").Handler(kitHttp.NewServer(
		endpoint,
		DecodeBasicRequest,
		EncodeBasicResponse,
		options...,
	))

	r.Methods("POST").Path("/sub/").Handler(kitHttp.NewServer(
		endpoint,
		DecodeBasicRequest,
		EncodeBasicResponse,
		options...,
	))

	r.Methods("POST").Path("/login/").Handler(kitHttp.NewServer(
		endpoint,
		DecodeBasicRequest,
		EncodeBasicResponse,
		options...,
	))

	return r
}

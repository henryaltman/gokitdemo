package transports

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"gokitdemo/core"
	"gokitdemo/util"
	"io/ioutil"
	"strings"
	//"log"
	"net/http"

	"gokitdemo/dto"

	//"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/endpoint"
	"github.com/gorilla/mux"

	"github.com/go-kit/kit/log"

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
func (tbr *TransportBaseReqquesst) DecodeGETRequest(r *http.Request) (interface{}, error) {
	//首字母设为大写
	path := strings.ReplaceAll(r.URL.Path, "/", "")
	if path != "" {
		path = util.Ucfirst(path)
	} else {
		path = "Default"
	}
	requestMethodName := fmt.Sprintf("Decode%sRequest", path)
	data := []byte{}
	if callResult := core.CallReflect(tbr, requestMethodName, data); callResult != nil {
		return callResult[0].Interface(), nil
	}
	return nil, errors.New("read body data error")
}

//DecodePOSTRequest is deal post request
func (tbr *TransportBaseReqquesst) DecodePOSTRequest(r *http.Request) (interface{}, error) {
	bodyBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, errors.New("read body data error")
	}
	//首字母设为大写
	path := strings.ReplaceAll(r.URL.Path, "/", "")
	if path != "" {
		path = util.Ucfirst(path)
	} else {
		path = "Default"
	}
	requestMethodName := fmt.Sprintf("Decode%sRequest", path)
	if callResult := core.CallReflect(tbr, requestMethodName, bodyBytes); callResult != nil {
		return callResult[0].Interface(), nil
	}
	return nil, errors.New("read body data error")
}

// decodeArithmeticRequest decode request params to struct
func decodeBasicRequest(_ context.Context, r *http.Request) (interface{}, error) {

	methodName := fmt.Sprintf("Decode%sRequest", r.Method)
	fmt.Println("methodName", methodName)
	tbr := &TransportBaseReqquesst{}
	if callResult := core.CallReflect(tbr, methodName, r); callResult != nil {
		callRet := callResult[0].Interface()
		fmt.Println(callRet)
		return callRet, nil
	}
	return nil, ErrorBadRequest
}

// encodeArithmeticResponse encode response to return
func encodeBasicResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

// MakeHttpHandler make http handler use mux
func MakeKitHttpHandler(ctx context.Context, endpoint endpoint.Endpoint, logger log.Logger) http.Handler {
	r := mux.NewRouter()
	options := []kitHttp.ServerOption{
		kitHttp.ServerErrorLogger(logger),
		kitHttp.ServerErrorEncoder(kitHttp.DefaultErrorEncoder),
	}

	r.Methods("GET").Path("/").Handler(kitHttp.NewServer(
		endpoint,
		decodeBasicRequest,
		encodeBasicResponse,
		options...,
	))

	r.Methods("POST").Path("/add/").Handler(kitHttp.NewServer(
		endpoint,
		decodeBasicRequest,
		encodeBasicResponse,
		options...,
	))

	r.Methods("POST").Path("/sub/").Handler(kitHttp.NewServer(
		endpoint,
		decodeBasicRequest,
		encodeBasicResponse,
		options...,
	))

	r.Methods("POST").Path("/login/").Handler(kitHttp.NewServer(
		endpoint,
		decodeBasicRequest,
		encodeBasicResponse,
		options...,
	))

	return r
}

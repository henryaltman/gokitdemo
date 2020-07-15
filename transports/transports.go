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

func (tbr *TransportBaseReqquesst) DecodeGETRequest(r *http.Request) (interface{}, error) {
	//首字母设为大写
	path := strings.ReplaceAll(r.URL.Path, "/", "")
	if path != "" {
		path = util.Ucfirst(path)
	} else {
		path = "Default"
	}
	fmt.Println("path", path)
	request := dto.BasicRequest{RequestId: "", Path: path, Data: ""}
	return request, nil
}

func (tbr *TransportBaseReqquesst) DecodePOSTRequest(r *http.Request) (interface{}, error) {
	bodyBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, errors.New("read body data error")
	}
	bodyMap := make(map[string]interface{})
	err = json.Unmarshal(bodyBytes, &bodyMap)
	if err != nil {
		return nil, errors.New("parse body data error")
	}
	//首字母设为大写
	path := strings.ReplaceAll(r.URL.Path, "/", "")
	if path != "" {
		path = util.Ucfirst(path)
	} else {
		path = "Default"
	}
	fmt.Println("path", path)
	request := dto.BasicRequest{RequestId: bodyMap["request_id"].(string), Path: path, Data: bodyMap["req"]}
	return request, nil
}

// decodeArithmeticRequest decode request params to struct
func decodeBasicRequest(_ context.Context, r *http.Request) (interface{}, error) {

	methodName := fmt.Sprintf("Decode%sRequest", r.Method)
	fmt.Println("methodName", methodName)
	tbr := &TransportBaseReqquesst{}
	if callResult := core.CallReflect(tbr, methodName, r); callResult != nil {
		return callResult[0].Interface(), nil
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
	rBytes, _ := json.Marshal(r)
	fmt.Println("r", string(rBytes))
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

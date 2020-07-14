package transports

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"

	//"log"
	"net/http"

	"gokitdemo/dto"

	//"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/endpoint"
	"github.com/gorilla/mux"

	"github.com/go-kit/kit/log"

	kitHttp "github.com/go-kit/kit/transport/http"
)

func encodeLoginResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

// decodeArithmeticRequest decode request params to struct
func decodeBasicRequest(_ context.Context, r *http.Request) (interface{}, error) {

	bodyBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("err", err)
		return nil, errors.New("read body data error")
	}
	fmt.Println("body", string(bodyBytes))
	bodyMap := make(map[string]interface{})
	err = json.Unmarshal(bodyBytes, &bodyMap)
	if err != nil {
		fmt.Println("err", err)
		return nil, errors.New("parse body data error")
	}

	if _, ok := bodyMap["request_id"]; !ok {
		return nil, errors.New("parse body request_id required error")
	}

	if _, ok := bodyMap["request_type"]; !ok {
		return nil, errors.New("parse body request_type required error")
	}

	request := dto.BasicRequest{RequestId: bodyMap["request_id"].(string), RequestType: bodyMap["request_type"].(string), Data: bodyMap["req"]}
	return request, nil
}

// encodeArithmeticResponse encode response to return
func encodeBasicResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

var (
	ErrorBadRequest = errors.New("invalid request parameter")
)

// MakeHttpHandler make http handler use mux
func MakeKitHttpHandler(ctx context.Context, endpoint endpoint.Endpoint, logger log.Logger) http.Handler {
	r := mux.NewRouter()
	rBytes, _ := json.Marshal(r)
	fmt.Println("r", string(rBytes))
	options := []kitHttp.ServerOption{
		kitHttp.ServerErrorLogger(logger),
		kitHttp.ServerErrorEncoder(kitHttp.DefaultErrorEncoder),
	}

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

	return r
}

package transports

import (
	"context"
	"encoding/json"
	"errors"

	//"log"
	"net/http"
	"strconv"

	"gokitdemo/endpoints"

	//"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/endpoint"
	"github.com/gorilla/mux"

	"github.com/go-kit/kit/log"

	kitHttp "github.com/go-kit/kit/transport/http"
)

// decodeArithmeticRequest decode request params to struct
func decodeBasicRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	requestType, ok := vars["type"]
	if !ok {
		return nil, ErrorBadRequest
	}

	pa, ok := vars["a"]
	if !ok {
		return nil, ErrorBadRequest
	}

	pb, ok := vars["b"]
	if !ok {
		return nil, ErrorBadRequest
	}

	a, _ := strconv.Atoi(pa)
	b, _ := strconv.Atoi(pb)

	return endpoints.BasicRequest{
		RequestType: requestType,
		A:           a,
		B:           b,
	}, nil
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

	options := []kitHttp.ServerOption{
		kitHttp.ServerErrorLogger(logger),
		kitHttp.ServerErrorEncoder(kitHttp.DefaultErrorEncoder),
	}

	r.Methods("POST").Path("/calculate/{type}/{a}/{b}").Handler(kitHttp.NewServer(
		endpoint,
		decodeBasicRequest,
		encodeBasicResponse,
		options...,
	))

	return r
}

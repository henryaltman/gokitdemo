package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"gokitdemo/endpoints"
	"gokitdemo/services"
	"gokitdemo/transports"

	"github.com/go-kit/kit/log"
)

func main() {

	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.With(logger, "ts", log.DefaultTimestampUTC)
		logger = log.With(logger, "caller", log.DefaultCaller)
	}
	ctx := context.Background()
	errChan := make(chan error)
	//var svc services.Service
	svcs := services.LoggingMiddleware(logger)

	endpoint := endpoints.MakeBasicEndpoint(svcs)

	r := transports.MakeKitHttpHandler(ctx, endpoint, logger)

	go func() {
		fmt.Println("Http Server start at port:9000")
		handler := r
		errChan <- http.ListenAndServe(":9000", handler)
	}()

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errChan <- fmt.Errorf("%s", <-c)
	}()

	fmt.Println(<-errChan)
}

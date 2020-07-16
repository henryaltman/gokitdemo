package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"gokitdemo/endpoints"
	"gokitdemo/rtlimit"
	"gokitdemo/services"
	"gokitdemo/transports"

	"github.com/go-kit/kit/log"
	"github.com/juju/ratelimit"
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
	// add ratelimit,refill every second,set capacity 3
	ratebucket := ratelimit.NewBucket(time.Second*1, 100000)
	endpoint = rtlimit.NewTokenBucketLimitterWithJuju(ratebucket)(endpoint)

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

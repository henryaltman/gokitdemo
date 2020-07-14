// loggingMiddleware Make a new type
// that contains Service interface and logger instance
package services

import (
	"context"
	"gokitdemo/dto"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/metrics"
)

type loggingMiddleware struct {
	Service
	logger log.Logger
}

type metricMiddleware struct {
	Service
	requestCount   metrics.Counter
	requestLatency metrics.Histogram
}

// func Metrics(requestCount metrics.Counter, requestLatency metrics.Histogram) ServiceMiddleware {
// 	return func(next Service) Service {
// 		return metricMiddleware{
// 			next,
// 			requestCount,
// 			requestLatency}
// 	}
// }

// LoggingMiddleware make logging middleware
func LoggingMiddleware(logger log.Logger) loggingMiddleware {
	svc := BasicService{}
	return loggingMiddleware{svc, logger}
}

func (mw loggingMiddleware) Add(ctx context.Context, r dto.AddRequest) (ret dto.AddResponse, err error) {

	// defer func(beign time.Time) {
	// 	lvs := []string{"method", "Add"}
	// 	mw.requestCount.With(lvs...).Add(1)
	// 	mw.requestLatency.With(lvs...).Observe(time.Since(beign).Seconds())
	// }(time.Now())

	ret, err = mw.Service.Add(ctx, r)
	return ret, err
}

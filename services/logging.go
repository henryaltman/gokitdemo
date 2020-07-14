// loggingMiddleware Make a new type
// that contains Service interface and logger instance
package services

import (
	"context"

	"github.com/go-kit/kit/log"
)

type loggingMiddleware struct {
	Service
	logger log.Logger
}

// LoggingMiddleware make logging middleware
func LoggingMiddleware(logger log.Logger) loggingMiddleware {
	svc := BasicService{}
	return loggingMiddleware{svc, logger}
}

func (mw loggingMiddleware) Add(ctx context.Context, r interface{}) BaseResponse {
	br := BaseResponse{}
	ret := mw.Service.Add(ctx, r)
	br.Rs = ret
	return br
}

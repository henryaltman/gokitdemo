// loggingMiddleware Make a new type
// that contains Service interface and logger instance
package services

import (
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

//func (mw loggingMiddleware) Add(ctx context.Context, r interface{}) BaseResponse {
//	br := BaseResponse{}
//	ret := mw.Service.Add(ctx, r)
//	br.Rs = ret
//	fmt.Println("br", br)
//	return br
//}
//
//func (mw loggingMiddleware) Subtract(ctx context.Context, r interface{}) BaseResponse {
//	br := BaseResponse{}
//	ret := mw.Service.Subtract(ctx, r)
//	br.Rs = ret
//	fmt.Println("br", br)
//	return br
//}

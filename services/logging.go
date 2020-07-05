// loggingMiddleware Make a new type
// that contains Service interface and logger instance
package services

import (
	"time"

	"github.com/go-kit/kit/log"
)

type loggingMiddleware struct {
	Service
	logger log.Logger
}

// func newService() BasicService {
// 	svc := BasicService{}
// 	return svc
// }

// LoggingMiddleware make logging middleware
func LoggingMiddleware(logger log.Logger) loggingMiddleware {
	svc := BasicService{}
	return loggingMiddleware{svc, logger}
}

func (mw loggingMiddleware) Add(a, b int) (ret int) {

	defer func(beign time.Time) {
		mw.logger.Log(
			"function", "Add",
			"a", a,
			"b", b,
			"result", ret,
			"took", time.Since(beign),
		)
	}(time.Now())

	ret = mw.Service.Add(a, b)
	return ret
}

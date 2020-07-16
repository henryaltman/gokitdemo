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

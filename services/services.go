package services

import (
	"context"
)

type BaseResponse struct {
	Rs  interface{}
	Err error
}

type (
	Service interface {

		// Add calculate a+b
		Add(context.Context, interface{}) BaseResponse

		// Subtract calculate a-b
		Subtract(a, b int) (int, error)

		// Multiply calculate a*b
		Multiply(a, b int) (int, error)

		// Divide calculate a/b
		Divide(a, b int) (int, error)

		Login(name, pwd string) (string, error)
	}
	BasicService struct{}
)

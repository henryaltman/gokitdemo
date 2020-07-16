package services

import (
	"context"
	"gokitdemo/dto"
)

type BaseResponse struct {
	Rs  interface{} `json:"rs"`
	Err error       `json:"err"`
}

type (
	Service interface {

		// Add calculate a+b
		Add(context.Context, dto.AddRequest) BaseResponse

		// Subtract calculate a-b
		Subtract(context.Context, interface{}) BaseResponse

		// Multiply calculate a*b
		Multiply(a, b int) (int, error)

		// Divide calculate a/b
		Divide(a, b int) (int, error)

		Login(context.Context, interface{}) string
		Default(context.Context, interface{}) BaseResponse
	}
	BasicService struct{}
)

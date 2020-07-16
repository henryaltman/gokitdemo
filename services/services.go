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
		//default router
		Default(context.Context, interface{}) BaseResponse
	}
	BasicService struct{}
)

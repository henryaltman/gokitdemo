package services

import (
	"context"
	"gokitdemo/dto"
)

type (
	Service interface {

		// Add calculate a+b
		Add(context.Context, dto.AddRequest) dto.BaseResponse
		//default router
		Default(context.Context, interface{}) dto.BaseResponse
	}
	BasicService struct{}
)

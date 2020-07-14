package services

import (
	"context"
	"gokitdemo/dto"
)

type (
	Service interface {

		// Add calculate a+b
		Add(context.Context, dto.AddRequest) (dto.AddResponse, error)

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

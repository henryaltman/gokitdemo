package services

import (
	"context"
	"gokitdemo/dto"
)

// Add implement Add method
func (s BasicService) Add(ctx context.Context, r dto.AddRequest) (rs dto.AddResponse, err error) {
	rs.Sum = r.A + r.B
	return rs, nil
}

package services

import (
	"context"
	"gokitdemo/dto"
)

// Add implement Add method
func (s BasicService) Add(ctx context.Context, r dto.AddRequest) (br dto.BaseResponse) {
	rs := dto.AddResponse{}
	rs.Sum = r.A + r.B
	br.Rs = rs
	return br
}

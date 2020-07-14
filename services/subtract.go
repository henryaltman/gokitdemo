package services

import (
	"context"
	"encoding/json"
	"fmt"
	"gokitdemo/dto"
)

// Subtract implement Subtract method
func (s BasicService) Subtract(ctx context.Context, r interface{}) BaseResponse {
	rs := dto.AddResponse{}
	var err error
	req := dto.AddRequest{}
	rByte, err := json.Marshal(r)
	br := BaseResponse{}
	if err != nil {
		br.Err = err
		return br
	}
	fmt.Println("rByte", string(rByte))
	err = json.Unmarshal(rByte, &req)

	rs.Sum = req.A + req.B

	br.Rs = rs
	rsByte, err := json.Marshal(rs)
	if err != nil {
		br.Err = err
		return br
	}
	fmt.Println("rsByte", string(rsByte))
	brByte, err := json.Marshal(br)
	if err != nil {
		br.Err = err
		return br
	}
	fmt.Println("brByte", string(brByte))
	return br
}

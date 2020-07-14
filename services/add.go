package services

import (
	"context"
	"encoding/json"
	"fmt"
	"gokitdemo/dto"
)

// Add implement Add method
func (s BasicService) Add(ctx context.Context, r interface{}) (br BaseResponse) {
	rs := dto.AddResponse{}
	var err error
	req := dto.AddRequest{}
	rByte, err := json.Marshal(r)
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
	return br
}

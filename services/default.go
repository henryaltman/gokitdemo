package services

import (
	"context"
	"gokitdemo/dto"
)

func (s BasicService) Default(ctx context.Context, r interface{}) (br dto.BaseResponse) {
	//tk := ctx.Value("tk")
	//fmt.Println("tk",tk)
	br.Rs = "hello world"
	return br
}

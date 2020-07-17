package services

import (
	"context"
)

func (s BasicService) Default(ctx context.Context, r interface{}) (br BaseResponse) {
	//tk := ctx.Value("tk")
	//fmt.Println("tk",tk)
	br.Rs = "hello world"
	return br
}

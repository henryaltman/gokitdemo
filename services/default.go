package services

import "context"

func (s BasicService) Default(ctx context.Context, r interface{}) (br BaseResponse) {
	br.Rs = "hello world"
	return br
}

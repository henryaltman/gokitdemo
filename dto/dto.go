package dto

type Base struct {
	UserId int
}
type BaseResponse struct {
	Rs  interface{} `json:"rs"`
	Err error       `json:"err"`
}

type BasicRequest struct {
	RequestId string
	Path      string //url path
	Request   interface{}
}

type BasicResponse struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

type AddRequest struct {
	Base
	A int `json:"a"`
	B int `json:"b"`
}

type AddResponse struct {
	Sum int `json:"sum"`
}

type DefaultRequest struct {
	Base
}

func (r DefaultRequest) Failed() error { return r.Failed() }

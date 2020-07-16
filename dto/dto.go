package dto

// {
// 	"request_type":"add",
// 	"RequestId":"xxxxx",
// 	"req_parameters":"{"a":1,"b":2}"
// }

type BasicRequest struct {
	RequestId string
	Path      string //url path
	Request   interface{}
}

type BasicResponse struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

type AddRequest struct {
	A int `json:"a"`
	B int `json:"b"`
}

type AddResponse struct {
	Sum int `json:"sum"`
}

type DefaultRequest struct {
}

type SubstractRequest struct {
	A int `json:"a"`
	B int `json:"b"`
}

type MultiplyRequest struct {
	A int `json:"a"`
	B int `json:"b"`
}

type DivideRequest struct {
	A int `json:"a"`
	B int `json:"b"`
}

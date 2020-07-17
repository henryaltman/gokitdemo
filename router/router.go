package router

//需要token

type R struct {
	VerifyToken bool
}

var Router = map[string]R{
	"Add": {VerifyToken: false},
}

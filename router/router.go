package router

//需要token

type R struct {
	Method      string
	VerifyToken bool
	Path        string
}

var Router = map[string]R{
	"Add":     {Path: "/add/", VerifyToken: false, Method: "POST"},
	"Default": {Path: "/", VerifyToken: false, Method: "GET"},
}

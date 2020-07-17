package router

//需要token
var RouterToken = map[string]bool{
	"Add": true,
}

//不需要token
var RouterWithoutToken = map[string]bool{
	"/": true,
}

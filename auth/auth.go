package auth

import (
	"context"
	"gokitdemo/util"
	stdHTTP "net/http"
	"strings"

	"github.com/go-kit/kit/transport/http"
)

const (
	Language        = "lang"
	LanguageHttpKey = "Content-Language"
	DefaultLanguage = "en"
	MethodKey       = "method"
	HttpPATH        = "path"
)

// HTTPToContext moves a tk from request header to context
func HTTPToContext() http.RequestFunc {
	return func(ctx context.Context, r *stdHTTP.Request) context.Context {
		//首字母设为大写
		path := strings.ReplaceAll(r.URL.Path, "/", "")
		if path != "" {
			path = util.Ucfirst(path)
		} else {
			path = "Default"
		}
		return context.WithValue(ctx, HttpPATH, path)
	}
}

// HTTPToContext moves a lang from request header to context
func LangHTTPToContext() http.RequestFunc {
	return func(ctx context.Context, r *stdHTTP.Request) context.Context {
		lang := r.Header.Get(LanguageHttpKey)
		if lang == "" {
			lang = DefaultLanguage
		}
		return context.WithValue(ctx, Language, lang)
	}
}

// HTTPToContext moves a lang from request header to context
func AuthorizationHTTPToContext() http.RequestFunc {
	return func(ctx context.Context, r *stdHTTP.Request) context.Context {
		return context.WithValue(ctx, MethodKey, r.Method)
	}
}

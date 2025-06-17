package router

import "net/http"

type Middleware func(http.Handler) http.Handler

type Router interface {
	Handle(method, pattern string, h http.HandlerFunc)
	Group(prefix string, fn func(r Router)) // optional
	Use(mw ...Middleware)                   // optional
	http.Handler                            // embeds ServeHTTP
}

type RouteInstaller func(Router)

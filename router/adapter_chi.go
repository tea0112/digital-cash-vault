package router

import (
	"github.com/go-chi/chi/v5"
	"net/http"
)

type ChiRouter struct{ chiRouter chi.Router }

func NewChi() Router { return &ChiRouter{chiRouter: chi.NewRouter()} }

func (r *ChiRouter) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	r.chiRouter.ServeHTTP(w, req)
}

func (r *ChiRouter) Handle(method, pattern string, h http.HandlerFunc) {
	r.chiRouter.MethodFunc(method, pattern, h)
}
func (r *ChiRouter) Group(prefix string, fn func(Router)) {
	r.chiRouter.Route(prefix, func(cr chi.Router) { fn(&ChiRouter{cr}) })
}
func (r *ChiRouter) Use(mw ...Middleware) {
	for _, m := range mw {
		r.chiRouter.Use(m)
	}
}

package router

import (
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"net/http"
)

type MuxRouter struct{ *mux.Router }

func NewMux() Router { return &MuxRouter{mux.NewRouter()} }

func (r *MuxRouter) Handle(method, pattern string, h http.HandlerFunc) {
	r.Router.HandleFunc(pattern, h).Methods(method)
}
func (r *MuxRouter) Group(prefix string, fn func(Router)) {
	sr := r.PathPrefix(prefix).Subrouter()
	fn(&MuxRouter{sr})
}
func (r *MuxRouter) Use(mw ...Middleware) {
	for _, m := range mw {
		r.Router.Use(func(next http.Handler) http.Handler { return m(next) })
	}
}

func muxRequestID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		req.Header.Set("X-Request-ID", uuid.NewString())
		next.ServeHTTP(w, req)
	})
}

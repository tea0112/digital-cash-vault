package router

import (
	"github.com/go-chi/chi/v5/middleware"
	"github.com/gorilla/handlers"
	"io"
	"net/http"
	"os"
)

func InstallDefaultMiddlewares(r Router) {
	switch v := r.(type) {
	case *ChiRouter:
		v.Use(middleware.RequestID, middleware.RealIP, middleware.Logger, middleware.Recoverer)

	case *MuxRouter:
		v.Use(muxRequestID, handlers.ProxyHeaders, Logging(os.Stdout))

	default:
		// fallback – maybe just logging/recovery implemented by you
	}
}

func Logging(out io.Writer) Middleware {
	return func(next http.Handler) http.Handler {
		return handlers.LoggingHandler(out, next)
	}
}

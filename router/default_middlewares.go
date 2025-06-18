package router

import (
	chimw "github.com/go-chi/chi/v5/middleware"
	"github.com/gorilla/handlers"
	"os"
)

func InstallDefaultMiddlewares(r Router) {
	switch v := r.(type) {
	case *ChiRouter:
		v.Use(chimw.RequestID, chimw.RealIP, chimw.Logger, chimw.Recoverer)

	case *MuxRouter:
		v.Use(muxRequestID, handlers.ProxyHeaders, Logging(os.Stdout))

	default:
		// fallback – maybe just logging/recovery implemented by you
	}
}

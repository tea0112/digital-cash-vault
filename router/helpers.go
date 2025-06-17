package router

import (
	"github.com/gorilla/handlers"
	"io"
	"net/http"
)

func Logging(out io.Writer) Middleware {
	return func(next http.Handler) http.Handler {
		return handlers.LoggingHandler(out, next)
	}
}

package router

import (
	"digital-cash-vault/pkg/customerrors"
	"encoding/json"
	"errors"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/gorilla/handlers"
	"io"
	"log"
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

func ErrorHandler(next func(w http.ResponseWriter, r *http.Request) error) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := next(w, r)
		if err != nil {
			// Log the internal error
			log.Printf("An error occurred: %v", err)

			var appErr *customerrors.AppError
			if as := AsAppError(err, &appErr); as {
				w.Header().Set("Content-Type", "application/json; charset=utf-8")
				w.WriteHeader(appErr.Code)
				EncodeErr := json.NewEncoder(w).Encode(appErr)
				if EncodeErr != nil {
					log.Printf(EncodeErr.Error())
					return
				}

				return
			}

			// Handle non-AppError errors
			w.Header().Set("Content-Type", "application/json; charset=utf-8")
			w.WriteHeader(http.StatusInternalServerError)
			EncodeErr := json.NewEncoder(w).Encode(map[string]string{"message": "Internal Server Error"})
			if EncodeErr != nil {
				log.Printf(EncodeErr.Error())
				return
			}
		}
	}
}

func AsAppError(err error, target **customerrors.AppError) bool {
	return errors.As(err, target)
}

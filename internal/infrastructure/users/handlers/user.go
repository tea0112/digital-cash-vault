package handlers

import (
	"digital-cash-vault/router"
	"log"
	"net/http"
)

func Register(r router.Router) {
	r.Group("/api", func(api router.Router) {
		api.Handle("GET", "/users", func(w http.ResponseWriter, r *http.Request) {
			log.Println("GET /users")
		})
		api.Handle("POST", "/users", func(w http.ResponseWriter, r *http.Request) {
			log.Println("POST /users")
		})
	})
}

func InitRouter(kind string) (router.Router, error) {
	r, err := router.New(kind)
	if err != nil {
		return nil, err
	}

	router.InstallDefaultMiddlewares(r)

	Register(r)
	return r, nil
}

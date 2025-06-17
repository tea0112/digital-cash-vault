package http

import (
	"digital-cash-vault/internal/interfaces/http/users"
	"digital-cash-vault/pkg/static"
	"digital-cash-vault/router"
)

func InitRouter(kind string, handlers map[string]any) (router.Router, error) {
	r, err := router.New(kind)
	if err != nil {
		return nil, err
	}

	router.InstallDefaultMiddlewares(r)

	users.RegisterRoutes(handlers[static.HandlerUser].(users.UserHandler))(r)

	return r, nil
}

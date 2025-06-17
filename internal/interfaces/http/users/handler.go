package users

import (
	"digital-cash-vault/internal/applications/users"
	"digital-cash-vault/pkg"
	"digital-cash-vault/router"
	"encoding/json"
	"net/http"
)

type UserHandler interface {
	GetUser(w http.ResponseWriter, r *http.Request)
	CreateUser(w http.ResponseWriter, r *http.Request)
	Login(w http.ResponseWriter, r *http.Request)
}

type UserHandlerImpl struct {
	userService users.UserService
}

func NewUserHandler(userService users.UserService) UserHandler {
	return &UserHandlerImpl{userService: userService}
}

func RegisterRoutes(userHandler UserHandler) router.RouteInstaller {
	return func(r router.Router) {
		r.Group("/api/v1", func(api router.Router) {
			api.Handle("GET", "/users", userHandler.GetUser)
			api.Handle("POST", "/users", userHandler.CreateUser)
			api.Handle("POST", "/users/login", userHandler.Login)
		})
	}
}

func (h *UserHandlerImpl) GetUser(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("Hello World"))
	if err != nil {
		return
	}
}

func (h *UserHandlerImpl) Login(w http.ResponseWriter, r *http.Request) {
	loginRequest, err := pkg.BindJSON[LoginRequestDTO](r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	response, err := h.userService.Login(r.Context(), loginRequest.ToInput())
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	res := LoginResponseDTO{
		User: UserDTO{
			ID:        string(response.User.Id),
			Username:  response.User.Username,
			Email:     response.User.Email,
			FirstName: response.User.FirstName,
			LastName:  response.User.LastName,
		},
		Token: response.Token,
	}

	w.Header().Set("Content-Type", "application/json")

	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func (h *UserHandlerImpl) CreateUser(w http.ResponseWriter, r *http.Request) {
	createUserRequestDTO, err := pkg.BindJSON[CreateUserRequestDTO](r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.userService.CreateUser(r.Context(), CreateUserRequestDTOToDomain(*createUserRequestDTO))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
}

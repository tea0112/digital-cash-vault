package users

import (
	"digital-cash-vault/internal/applications/users"
	"digital-cash-vault/pkg"
	"digital-cash-vault/pkg/customerrors"
	"digital-cash-vault/router"
	"encoding/json"
	"net/http"
)

type UserHandler interface {
	GetUser(w http.ResponseWriter, r *http.Request) error
	CreateUser(w http.ResponseWriter, r *http.Request) error
	Login(w http.ResponseWriter, r *http.Request) error
	GetList(w http.ResponseWriter, r *http.Request) error
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
			api.Handle("GET", "/users", router.ErrorHandler(userHandler.GetUser))
			api.Handle("POST", "/users", router.ErrorHandler(userHandler.CreateUser))
			api.Handle("POST", "/users/login", router.ErrorHandler(userHandler.Login))
			api.Handle("GET", "/users", router.ErrorHandler(userHandler.GetList))
		})
	}
}

func (h *UserHandlerImpl) GetUser(w http.ResponseWriter, r *http.Request) error {
	_, err := w.Write([]byte("Hello World"))
	if err != nil {
		return customerrors.NewAppError(http.StatusNotFound, "User not found", err)
	}

	return nil
}

func (h *UserHandlerImpl) Login(w http.ResponseWriter, r *http.Request) error {
	loginRequest, err := pkg.BindJSON[LoginRequestDTO](r)
	if err != nil {
		return customerrors.NewAppError(http.StatusBadRequest, "Login payload wrong", err)
	}

	response, err := h.userService.Login(r.Context(), loginRequest.ToInput())
	if err != nil {
		return customerrors.NewAppError(http.StatusBadRequest, "User's credentials wrong", err)
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
		return customerrors.NewAppError(http.StatusInternalServerError, "Server Internal Error", err)
	}

	return nil
}

func (h *UserHandlerImpl) CreateUser(w http.ResponseWriter, r *http.Request) error {
	createUserRequestDTO, err := pkg.BindJSON[CreateUserRequestDTO](r)
	if err != nil {
		return customerrors.NewAppError(http.StatusInternalServerError, "Server Internal Error", err)
	}

	err = h.userService.CreateUser(r.Context(), CreateUserRequestDTOToDomain(*createUserRequestDTO))
	if err != nil {
		return customerrors.NewAppError(http.StatusBadRequest, "Cannot create user", err)
	}

	return nil
}

func (h *UserHandlerImpl) GetList(w http.ResponseWriter, r *http.Request) error {
	query := r.URL.Query()

	// Parse pagination
	page := pkg.ParseInt(query.Get("page"), 1)
	size := pkg.ParseInt(query.Get("size"), 20)

	// Optional filters
	var usernameLike *string
	if val := query.Get("username"); val != "" {
		usernameLike = &val
	}

	var status *string
	if val := query.Get("status"); val != "" {
		status = &val
	}

	var emailLike *string
	if val := query.Get("email"); val != "" {
		emailLike = &val
	}

	params := users.ListUsersParams{
		Username:  usernameLike,
		Status:    status,
		EmailLike: emailLike,
		OrderBy:   "created_at DESC",
		Page:      page,
		Size:      size,
	}

	usersList, err := h.userService.ListUsers(r.Context(), params)
	if err != nil {
		return customerrors.NewAppError(http.StatusBadRequest, "Cannot get user list", err)
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(FromDomainList(usersList))
	if err != nil {
		return customerrors.NewAppError(http.StatusInternalServerError, "Server Internal Error", err)
	}

	return nil
}

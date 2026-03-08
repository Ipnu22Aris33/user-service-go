package user

import (
	"net/http"

	"github.com/Ipnu22Aris33/user-service-go/internal/httpx"
)

type UserHandler struct {
	usecase *UserUsecase
}

func NewUserHandler(u *UserUsecase) *UserHandler {
	return &UserHandler{
		usecase: u,
	}
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var req CreateUserRequest

	if err := httpx.DecodeAndValidate(r, &req); err != nil {
		httpx.Error(w, http.StatusBadRequest, err)
		return
	}

	user := &User{
		Email: req.Email,
		Name:  req.Name,
	}

	user, err := h.usecase.CreateUser(user)

	if err != nil {
		httpx.Error(w, http.StatusInternalServerError, err)
		return
	}

	httpx.JSON(w, http.StatusCreated, user)
}

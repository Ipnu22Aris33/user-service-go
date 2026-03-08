package user

import (
	"github.com/Ipnu22Aris33/user-service-go/internal/httpx"
	"github.com/go-chi/chi/v5"
	"net/http"
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

func (h *UserHandler) GetUserByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	user, err := h.usecase.GetUserByID(id)
	if err != nil {
		httpx.Error(w, http.StatusInternalServerError, err)
		return
	}

	httpx.JSON(w, http.StatusOK, user)
}

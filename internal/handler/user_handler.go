package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/Ipnu22Aris33/user-service-go/internal/domain"
	"github.com/Ipnu22Aris33/user-service-go/internal/usecase"
)

type UserHandler struct {
	usecase *usecase.UserUsecase
}

func NewUserHandler(u *usecase.UserUsecase) *UserHandler {
	return &UserHandler{
		usecase: u,
	}
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var user domain.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err := h.usecase.CreateUser(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, user)
}
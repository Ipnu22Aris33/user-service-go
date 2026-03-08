package server

import (
	"github.com/Ipnu22Aris33/user-service-go/internal/user"
)

type Container struct {
	UserHandler *user.UserHandler
}

func NewContainer() *Container {

	userRepo := user.NewInMemoryUserRepository()

	userUsecase := user.NewUserUsecase(userRepo)

	userHandler := user.NewUserHandler(userUsecase)

	return &Container{
		UserHandler: userHandler,
	}
}

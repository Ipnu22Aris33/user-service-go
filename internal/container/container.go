package container

import (
	"github.com/Ipnu22Aris33/user-service-go/internal/handler"
	"github.com/Ipnu22Aris33/user-service-go/internal/repository"
	"github.com/Ipnu22Aris33/user-service-go/internal/usecase"
)

type Container struct {
	UserHandler *handler.UserHandler
}

func NewContainer() *Container {

	userRepo := repository.NewInMemoryUserRepository()

	userUsecase := usecase.NewUserUsecase(userRepo)

	userHandler := handler.NewUserHandler(userUsecase)

	return &Container{
		UserHandler: userHandler,
	}
}

package usecase

import "github.com/Ipnu22Aris33/user-service-go/internal/domain"

type UserUsecase struct {
	repo domain.UserRepository
}

func NewUserUsecase(r domain.UserRepository) *UserUsecase {
	return &UserUsecase{
		repo: r,
	}
}

func (u *UserUsecase) CreateUser(user *domain.User) error {
	return u.repo.CreateUser(user)
}
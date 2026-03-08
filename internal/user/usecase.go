package user

import "github.com/google/uuid"

type UserUsecase struct {
	repo UserRepository
}

func NewUserUsecase(r UserRepository) *UserUsecase {
	return &UserUsecase{
		repo: r,
	}
}

func (u *UserUsecase) CreateUser(user *User) (*User, error) {
	user.ID = uuid.NewString()

	if err := u.repo.CreateUser(user); err != nil {
		return nil, err
	}

	return user, nil
}

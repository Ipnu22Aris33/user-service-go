package user

type UserUsecase struct {
	repo UserRepository
}

func NewUserUsecase(r UserRepository) *UserUsecase {
	return &UserUsecase{
		repo: r,
	}
}

func (u *UserUsecase) CreateUser(user *User) error {
	return u.repo.CreateUser(user)
}

package user

type User struct {
	ID    string `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`
}

type UserRepository interface {
	GetUserByID(id string) (*User, error)
	CreateUser(user *User) error
}

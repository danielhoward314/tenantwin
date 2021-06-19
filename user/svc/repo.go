package svc

type UserRepository interface {
	Signup(user *User) error
}

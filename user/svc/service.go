package svc

import "time"

type UserService interface {
	Signup(user *User) error
}

type userServiceImpl struct {
	userRepo UserRepository
}

func NewUserService(userRepo UserRepository) UserService {
	return &userServiceImpl{
		userRepo: userRepo,
	}
}

func (us *userServiceImpl) Signup(user *User) error {
	user.CreatedAt = time.Now()
	return us.userRepo.Signup(user)
}

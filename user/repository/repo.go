package repository

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/danielhoward314/tenantwin/user/svc"
	"github.com/pkg/errors"
)

type userRepositoryImpl struct {
	client *gorm.DB
}

func NewUserRepository(DbHost, DbUser, DbPassword, DbName, DbPort string) (svc.UserRepository, error) {
	url := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", DbHost, DbUser, DbPassword, DbName, DbPort)
	client, err := gorm.Open(postgres.Open(url), &gorm.Config{})
	if err != nil {
		return nil, errors.Wrap(err, "userService.NewUserRepository")
	}
	// client.Debug().AutoMigrate(&svc.User{})
	repo := &userRepositoryImpl{client: client}
	return repo, nil
}

func (r *userRepositoryImpl) Signup(user *svc.User) error {
	err := r.client.Create(user).Error
	if err != nil {
		return err
	}
	return nil
}

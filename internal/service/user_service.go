package service

import (
	"github.com/AvinFajarF/internal/entity"
	"github.com/AvinFajarF/internal/repository"
)

type UserService interface {
	Register(username, password, email, image, bio string) (*entity.UserEntity, error)
	Login(username, password string) error
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repository repository.UserRepository) userService {
	return userService{
        repo: repository,
    }
}


package service

import (
	"github.com/AvinFajarF/internal/entity"
	"github.com/AvinFajarF/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	Register(username, password, email, image, bio string) (*entity.Users, error)
	Login(username, password string)  error
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repository repository.UserRepository) userService {
	return userService{
        repo: repository,
    }
}

func (s *userService) Register(username, password, email, image, bio string) (*entity.Users, error){
	user := &entity.Users{
		Username: username,
        Password: password,
        Email: email,
        Image: image,
        Bio: bio,
	}

	if err := s.repo.Save(user); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *userService) Login(email, password string)  error {
	user , err := s.repo.SigIn(email)

	if err != nil {
        return err
    }
	
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	if err != nil {
		return err
	}

	return nil
}
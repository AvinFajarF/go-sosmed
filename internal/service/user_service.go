package service

import	"github.com/AvinFajarF/internal/entity"

type UserService interface {
	Register(username, password, email, image, bio string) (*entity.UserEntity, error)
}
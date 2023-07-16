package repository

import (
	"github.com/AvinFajarF/internal/entity"
	"gorm.io/gorm"
)

type UserRepository interface {
	Save(user *entity.UserEntity) error
	SigIn(email , password string) (*entity.UserEntity, error)
}

type PostgreUserRepository struct {
	db *gorm.DB
}

func NewPostgreUserRepository(db *gorm.DB) *PostgreUserRepository {
	return &PostgreUserRepository{
        db: db,
    }
}

// function ini untuk register repository
func (r *PostgreUserRepository) Save(user *entity.UserEntity) error {
    return r.db.Create(user).Error
}

// function ini untuk login repository
func (r *PostgreUserRepository) SigIn(email, password string) (*entity.UserEntity, error) {
	var user entity.UserEntity 
	if err := r.db.Where("email =? and password =?", email, password).First(&user).Error ; err != nil {
		return nil, err
	}
	return &user, nil
}
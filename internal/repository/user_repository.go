package repository

import (
	"github.com/AvinFajarF/internal/entity"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRepository interface {
	Save(user *entity.Users) error
	SigIn(email  string) (*entity.Users, error)
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
func (r *PostgreUserRepository) Save(user *entity.Users) error {
	user.ID = uuid.NewString()
    return r.db.Create(user).Error
}

// function ini untuk login repository
func (r *PostgreUserRepository) SigIn(email string) (*entity.Users, error) {
	var user entity.Users 
	if err := r.db.Where("email = ? ", email).First(&user).Error ; err != nil {
		return nil, err
	}
	return &user, nil
}
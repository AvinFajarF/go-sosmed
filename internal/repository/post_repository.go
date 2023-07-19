package repository

import (
	"github.com/AvinFajarF/internal/entity"
	"gorm.io/gorm"
)

type PostRepository interface {
	Create(post *entity.Posts) error
}

type PostgrePostssRepository struct {
	db *gorm.DB
}

func NewPostgrePostRepository(db *gorm.DB) *PostgrePostssRepository {
	return &PostgrePostssRepository{
		db: db,
	}
}

func (r *PostgrePostssRepository) Create(post *entity.Posts) error {
    return r.db.Create(post).Error
}
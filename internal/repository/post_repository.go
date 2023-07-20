package repository

import (
	"github.com/AvinFajarF/internal/entity"
	"gorm.io/gorm"
)

type PostRepository interface {
	Create(post *entity.Posts) error
	Get() (post *entity.Posts, err error)
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

func (r *PostgrePostssRepository) Get() (post *entity.Posts, err error) {
    err = r.db.Find(&post).Error
    return
}
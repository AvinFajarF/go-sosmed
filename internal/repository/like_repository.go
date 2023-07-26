package repository

import (
	"github.com/AvinFajarF/internal/entity"
	"github.com/google/uuid"
)

type LikeRepository interface {
	Likes(data *entity.Likes) error
}

func (r *PostgrePostssRepository) Like(data *entity.Likes) error {
	data.ID = uuid.NewString()
	return r.db.Create(data).Error
}
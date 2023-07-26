package repository

import (
	"log"

	"github.com/AvinFajarF/internal/entity"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PostRepository interface {
	Create(post *entity.Posts) error
	FindUserId(userId string, c *gin.Context) error
	Get(userId string, c *gin.Context) (post []entity.Posts, err error)
	Delete(id string) error
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
	post.ID = uuid.NewString()
	return r.db.Create(post).Error
}

func (r *PostgrePostssRepository) FindUserId(userId string, c *gin.Context) error {
	log.Println(userId)
	return r.db.Where("user_id = ? ", userId).Find(&entity.Posts{}).Error
}

func (r *PostgrePostssRepository) Get(userId string, c *gin.Context) (post []entity.Posts, err error) {
	err = r.db.Where("user_id  = ?", userId).Find(&post).Error
	return 
}

func (r *PostgrePostssRepository) Delete(id string) error {
	return r.db.Where("id = ?", id).Delete(&entity.Posts{}).Error
}

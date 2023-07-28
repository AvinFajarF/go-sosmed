package repository

import (
	"fmt"

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
	Update(id string, data *entity.Posts) error
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
	  // Retrieve the user information from the database
	  var user entity.Users
	  if err := r.db.First(&user, "id = ?", post.UserID).Error; err != nil {
		  return err
	  }
  
	  // Set the User field in the post struct to the retrieved user
	  post.User = user
  
	  // Create the post and save the association in the database
	  return r.db.Create(post).Error
}

func (r *PostgrePostssRepository) FindUserId(userId string, c *gin.Context) error {
	return r.db.Where("user_id = ? ", userId).Find(&entity.Posts{}).Error
}

func (r *PostgrePostssRepository) Get(userId string, c *gin.Context) (posts []entity.Posts, err error) {
    err = r.db.Joins("JOIN users ON posts.user_id = users.id").Where("user_id = ?", userId).Find(&posts).Error
    if err != nil {
        fmt.Printf("Error executing query: %s\n", err.Error())
        return nil, err
    }

    // Establish the association between Posts and Users
    for i := range posts {
        r.db.Model(&posts[i]).Association("User").Find(&posts[i].User)
    }

    return posts, nil
}



func (r *PostgrePostssRepository) Delete(id string) error {
	return r.db.Where("id = ?", id).Delete(&entity.Posts{}).Error
}

func (r *PostgrePostssRepository) Update(id string, data *entity.Posts) error {
	return r.db.Where("id = ?", id).Updates(data).Error
}
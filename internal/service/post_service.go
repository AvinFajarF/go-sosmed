package service

import (
	"github.com/AvinFajarF/internal/entity"
	"github.com/AvinFajarF/internal/repository"
	"github.com/gin-gonic/gin"
)

type PostService interface {
	CreatePosts(title, description, user_id string) (*entity.Posts, error)
	FindUserById(userId string, c *gin.Context) error
	GetPosts(userId string, c *gin.Context) ([]entity.Posts, error)
	DeletePost(id string) error
	UpdatePost(id string, title, description string) error
}

type postService struct {
	repo repository.PostRepository
}

func NewPostService(repo repository.PostRepository) postService {
	return postService{repo: repo}
}

func (ps postService) CreatePosts(title, description, user_id string) (*entity.Posts, error) {

	post := &entity.Posts{
		Title:       title,
		Description: description,
		UserID:      user_id,
	}

	if err := ps.repo.Create(post); err != nil {
		return nil, err
	}

	return post, nil
}

func (ps postService) GetPosts(userId string, c *gin.Context) ([]entity.Posts, error) {
	result, err := ps.repo.Get(userId, c)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (ps postService) DeletePost(id string) error {

	err := ps.repo.Delete(id)

	if err != nil {
		return err
	}

	return nil
}

func (ps postService) FindUserById(userId string, c *gin.Context) error {
	err := ps.repo.FindUserId(userId, c)

	if err != nil {
		return err
	}

	return nil
}

func (ps postService) UpdatePost(id string, title, description string) error {

	post := &entity.Posts{
		ID:          id,
		Title:       title,
		Description: description,
	}

	if err := ps.repo.Update(id, post); err != nil {
		return err
	}
	return nil
}

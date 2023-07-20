package service

import (
	"github.com/AvinFajarF/internal/entity"
	"github.com/AvinFajarF/internal/repository"
)

type PostService interface {
	CreatePosts(title, description, user_id string) (*entity.Posts, error)
	GetPosts() (*entity.Posts, error)
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

func (ps postService) GetPosts() (*entity.Posts, error) {
	result, err := ps.repo.Get()

	if err != nil {
		return nil, err
	}

	return result, nil
}

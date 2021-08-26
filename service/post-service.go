package service

import (
	"errors"
	"math/rand"
	"strconv"

	"github.com/underscoreanuj/mux_api/entity"
	"github.com/underscoreanuj/mux_api/repository"
)

type PostService interface {
	Validate(post *entity.Post) error
	Create(post *entity.Post) (*entity.Post, error)
	FindAll() ([]entity.Post, error)
	FindById(id string) (*entity.Post, error)
}

type service struct{}

var (
	postRepo repository.PostRepository
)

func NewPostService(repo repository.PostRepository) PostService {
	postRepo = repo
	return &service{}
}

func (*service) Validate(post *entity.Post) error {
	if post == nil {
		err := errors.New("empty post!")
		return err
	}
	if post.Title == "" {
		err := errors.New("post title is empty!")
		return err
	}

	return nil
}

func (*service) Create(post *entity.Post) (*entity.Post, error) {
	post.Id = rand.Int63()
	return postRepo.Save(post)
}

func (*service) FindAll() ([]entity.Post, error) {
	return postRepo.FindAll()
}

func (*service) FindById(id string) (*entity.Post, error) {
	_, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return nil, err
	}

	return postRepo.FindById(id)
}

package repository

import (
	"github.com/underscoreanuj/mux_api/entity"
)

type PostRepository interface {
	Save(post *entity.Post) (*entity.Post, error)
	FindAll() ([]entity.Post, error)
	Delete(post *entity.Post) error
}

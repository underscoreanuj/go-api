package cache

import "github.com/underscoreanuj/mux_api/entity"

type PostCache interface {
	Set(key string, value *entity.Post)
	Get(key string) *entity.Post
}

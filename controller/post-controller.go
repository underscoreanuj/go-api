package controller

import (
	"strings"

	"github.com/underscoreanuj/mux_api/cache"
	"github.com/underscoreanuj/mux_api/entity"
	"github.com/underscoreanuj/mux_api/errors"
	"github.com/underscoreanuj/mux_api/service"

	"encoding/json"
	"net/http"
)

type controller struct{}

var (
	postService service.PostService
	postCache   cache.PostCache
)

type PostController interface {
	GetPostById(response http.ResponseWriter, request *http.Request)
	GetPosts(resp http.ResponseWriter, req *http.Request)
	AddPost(resp http.ResponseWriter, req *http.Request)
}

func NewPostController(service service.PostService, cache cache.PostCache) PostController {
	postService = service
	postCache = cache
	return &controller{}
}

func (*controller) GetPostById(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-Type", "application/json")
	postId := strings.Split(req.URL.Path, "/")[2]
	var post *entity.Post = postCache.Get(postId)

	if post == nil {
		post, err := postService.FindById(postId)
		if err != nil {
			resp.WriteHeader(http.StatusNotFound)
			json.NewEncoder(resp).Encode(errors.ServiceError{Message: "No post found!"})
			return
		}
		postCache.Set(postId, post)

		resp.WriteHeader(http.StatusOK)
		json.NewEncoder(resp).Encode(post)

	} else {
		resp.WriteHeader(http.StatusOK)
		json.NewEncoder(resp).Encode(post)

	}
}

func (*controller) GetPosts(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-type", "application/json")

	posts, err := postService.FindAll()
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(resp).Encode(errors.ServiceError{Message: err.Error()})
		return
	}

	resp.WriteHeader(http.StatusOK)
	json.NewEncoder(resp).Encode(posts)
}

func (*controller) AddPost(resp http.ResponseWriter, req *http.Request) {
	var post entity.Post
	err := json.NewDecoder(req.Body).Decode(&post)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(resp).Encode(errors.ServiceError{Message: "Error unmarshalling data"})
		return
	}

	err1 := postService.Validate(&post)
	if err1 != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(resp).Encode(errors.ServiceError{Message: err1.Error()})
		return
	}

	result, err2 := postService.Create(&post)

	if err2 != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(resp).Encode(errors.ServiceError{Message: err2.Error()})
		return
	}

	resp.WriteHeader(http.StatusOK)
	json.NewEncoder(resp).Encode(result)
}

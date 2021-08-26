package controller

import (
	"github.com/underscoreanuj/mux_api/entity"
	"github.com/underscoreanuj/mux_api/errors"
	"github.com/underscoreanuj/mux_api/service"

	"encoding/json"
	"net/http"
)

type controller struct{}

var (
	postService service.PostService
)

type PostController interface {
	GetPosts(resp http.ResponseWriter, req *http.Request)
	AddPost(resp http.ResponseWriter, req *http.Request)
}

func NewPostController(service service.PostService) PostController {
	postService = service
	return &controller{}
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

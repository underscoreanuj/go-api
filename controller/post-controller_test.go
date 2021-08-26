package controller

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/underscoreanuj/mux_api/cache"
	"github.com/underscoreanuj/mux_api/entity"
	"github.com/underscoreanuj/mux_api/repository"
	"github.com/underscoreanuj/mux_api/service"
)

const (
	ID    int64  = 123
	TITLE string = "test title"
	TEXT  string = "test text"
)

var (
	postRepo       repository.PostRepository = repository.NewSQLiteRepository()
	postSrv        service.PostService       = service.NewPostService(postRepo)
	postCacheSrv   cache.PostCache           = cache.NewRedisCache("localhost:6379", 0, 10)
	postController PostController            = NewPostController(postSrv, postCacheSrv)
)

func setup() {
	var post entity.Post = entity.Post{
		Id:    ID,
		Title: TITLE,
		Text:  TEXT,
	}

	postRepo.Save(&post)
}

func cleanUp(post *entity.Post) {
	postRepo.Delete(post)
}

func TestAddPost(t *testing.T) {
	// create a new http post request
	var jsonReq = []byte(`{"title": "` + TITLE + `", "text": "` + TEXT + `"}`)
	req, _ := http.NewRequest("POST", "/posts", bytes.NewBuffer(jsonReq))

	// assign http handler function (controller AddPost function)
	handler := http.HandlerFunc(postController.AddPost)

	// record http response (httptest)
	response := httptest.NewRecorder()

	// dispatch the http request
	handler.ServeHTTP(response, req)

	// add assertions on the http status codes and the response

	status := response.Code

	if status != http.StatusOK {
		t.Errorf("Handler returned a wrong status code: got %v but expected %v", status, http.StatusOK)
	}

	// Decode the HTTP response
	var post entity.Post
	json.NewDecoder(io.Reader(response.Body)).Decode(&post)

	// assert HTTP response
	assert.NotNil(t, post.Id)
	assert.Equal(t, TITLE, post.Title)
	assert.Equal(t, TEXT, post.Text)

	// clean database
	cleanUp(&post)
}

func TestGetPosts(t *testing.T) {
	// add a post to the repo
	setup()

	// create a GET HTTP request
	req, _ := http.NewRequest("GET", "/posts", nil)

	// assign http handler function (controller GetPost function)
	handler := http.HandlerFunc(postController.GetPosts)

	// record http response (httptest)
	response := httptest.NewRecorder()

	// dispatch the http request
	handler.ServeHTTP(response, req)

	// add assertions on the http status codes and the response

	status := response.Code

	if status != http.StatusOK {
		t.Errorf("Handler returned a wrong status code: got %v but expected %v", status, http.StatusOK)
	}

	// Decode the HTTP response
	var posts []entity.Post
	json.NewDecoder(io.Reader(response.Body)).Decode(&posts)

	// assert HTTP response
	assert.NotNil(t, posts[0].Id)
	assert.Equal(t, TITLE, posts[0].Title)
	assert.Equal(t, TEXT, posts[0].Text)

	// clean database
	cleanUp(&posts[0])

}

func TestGetPostById(t *testing.T) {
	setup()

	// create a GET HTTP request
	req, _ := http.NewRequest("GET", "/posts/"+strconv.FormatInt(ID, 10), nil)

	// assign http handler function (controller GetPost function)
	handler := http.HandlerFunc(postController.GetPostById)

	// record http response (httptest)
	response := httptest.NewRecorder()

	// dispatch the http request
	handler.ServeHTTP(response, req)

	// add assertions on the http status codes and the response

	status := response.Code

	if status != http.StatusOK {
		t.Errorf("Handler returned a wrong status code: got %v but expected %v", status, http.StatusOK)
	}

	// Decode the HTTP response
	var posts entity.Post
	json.NewDecoder(io.Reader(response.Body)).Decode(&posts)

	// assert HTTP response
	assert.NotNil(t, posts.Id)
	assert.Equal(t, TITLE, posts.Title)
	assert.Equal(t, TEXT, posts.Text)

	// clean database
	cleanUp(&posts)

}

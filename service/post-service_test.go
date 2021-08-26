package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/underscoreanuj/mux_api/entity"
)

type MockRepository struct {
	mock.Mock
}

func (mock *MockRepository) Save(post *entity.Post) (*entity.Post, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.(*entity.Post), args.Error(1)
}

func (mock *MockRepository) FindAll() ([]entity.Post, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.([]entity.Post), args.Error(1)

}

func (mock *MockRepository) Delete(post *entity.Post) error {
	args := mock.Called()
	_ = args.Get(0)
	return args.Error(1)
}

func TestValidateEmptyPost(t *testing.T) {
	testService := NewPostService(nil)

	err := testService.Validate(nil)

	assert.NotNil(t, err)
	assert.Equal(t, "empty post!", err.Error())
}

func TestValidateEmptyPostTitle(t *testing.T) {
	post := entity.Post{Id: 1, Title: "", Text: "test text"}

	testService := NewPostService(nil)

	err := testService.Validate(&post)

	assert.NotNil(t, err)
	assert.Equal(t, "post title is empty!", err.Error())
}

func TestFindAll(t *testing.T) {
	mockRepo := new(MockRepository)
	var identifier int64 = 1
	post := entity.Post{Id: identifier, Title: "A", Text: "B"}

	mockRepo.On("FindAll").Return([]entity.Post{post}, nil)

	testService := NewPostService(mockRepo)

	result, err := testService.FindAll()

	mockRepo.AssertExpectations(t)
	assert.Nil(t, err)
	assert.Equal(t, identifier, result[0].Id)
	assert.Equal(t, "A", result[0].Title)
	assert.Equal(t, "B", result[0].Text)
}

func TestCreate(t *testing.T) {
	mockRepo := new(MockRepository)
	post := entity.Post{Title: "A", Text: "B"}

	mockRepo.On("Save").Return(&post, nil)

	testService := NewPostService(mockRepo)

	result, err := testService.Create(&post)

	mockRepo.AssertExpectations(t)
	assert.Nil(t, err)
	assert.NotNil(t, result.Id)
	assert.Equal(t, "A", result.Title)
	assert.Equal(t, "B", result.Text)
}

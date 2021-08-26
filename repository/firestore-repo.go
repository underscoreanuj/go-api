package repository

import (
	"context"
	"errors"
	"log"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"

	"github.com/underscoreanuj/mux_api/entity"
)

type repo struct{}

func NewFirestoreRepository() PostRepository {
	return &repo{}
}

const (
	projectId      string = "go-mux-api-proj"
	collectionName string = "posts"
)

func (*repo) Save(post *entity.Post) (*entity.Post, error) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectId)
	if err != nil {
		log.Fatalf("Failed to create a firestore client: %v", err)
		return nil, err
	}

	defer client.Close()

	_, _, err = client.Collection(collectionName).Add(ctx, map[string]interface{}{
		"Id":    post.Id,
		"Title": post.Title,
		"Text":  post.Text,
	})

	if err != nil {
		log.Fatalf("Failed adding a new post: %v", err)
		return nil, err
	}

	return post, nil
}

func (*repo) FindAll() ([]entity.Post, error) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectId)
	if err != nil {
		log.Fatalf("Failed to create a firestore client: %v", err)
		return nil, errors.New("Failed to create a firestore client")
	}

	defer client.Close()

	var posts []entity.Post
	it := client.Collection(collectionName).Documents(ctx)

	for {
		doc, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Failed to iterate over the posts: %v", err)
			return nil, errors.New("Failed to iterate over the posts")
		}

		post := entity.Post{
			Id:    doc.Data()["Id"].(int64),
			Title: doc.Data()["Title"].(string),
			Text:  doc.Data()["Text"].(string),
		}

		posts = append(posts, post)
	}

	if len(posts) == 0 {
		return nil, errors.New("No document found")
	}

	return posts, nil
}

// TODO
func (*repo) FindById(id string) (*entity.Post, error) {
	return nil, nil
}

// TODO
func (*repo) Delete(post *entity.Post) error {
	return nil
}

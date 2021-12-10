package repository

import (
	"context"
	"log"
	"main/src/entity"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
)

type repo struct {}
// NewFirestoreRepository
func NewFirestoreRepository() PostRepository {
	return &repo{}
}

const (
	projectID = "mux-api-98344"
	collectionName = "posts"
)

func (*repo) Save(post *entity.Post) (*entity.Post, error) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectID)

	if err != nil {
		log.Fatalf("Failed to create a firestore client: %v", err)
		return nil, err
	}

	defer client.Close()

	_, _, err = client.Collection(collectionName).Add(ctx, map[string]interface{}{
		"ID": post.ID,
		"Title": post.Title,
		"Text": post.Text,
	})

	if err != nil {
		log.Fatalf("Failed adding a new post: %v", err)
		return nil, err	
	}

	return post, nil
}

func (*repo) FindAll() ([]entity.Post, error) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectID)

	if err != nil {
		log.Fatalf("Failed to create a firestore client: %v", err)
		return nil, err
	}

	defer client.Close()

	var posts []entity.Post
	itr := client.Collection(collectionName).Documents(ctx)
	for {
		doc, err := itr.Next()

		if err == iterator.Done {
			break
		}

		if err != nil {
			log.Fatalf("Failed to iterate the list of posts: %v", err)
			return nil, err
		}

		post := entity.Post{
			ID: doc.Data()["ID"].(int64),
			Title: doc.Data()["Title"].(string),
			Text: doc.Data()["Text"].(string),
		}
		posts = append(posts, post)
	}

	return posts, nil

}
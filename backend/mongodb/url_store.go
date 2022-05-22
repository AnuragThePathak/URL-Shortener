package mongodb

import (
	"context"
	"errors"
	"time"

	"github.com/AnuragThePathak/url-shortener/backend/service"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const createIndexTimeout = 5 * time.Second

type urlStore struct {
	collection *mongo.Collection
}

func NewUrlStore(database *mongo.Database) (service.UrlStore, error) {
	context, cancel :=
		context.WithTimeout(context.Background(), createIndexTimeout)
	defer cancel()
	isUnique := true
	collection := database.Collection("urls")
	collection.Indexes().CreateMany(
		context, []mongo.IndexModel{
			{
				Keys: bson.D{{"original", 1}, {"shortened", 1}},
				Options: &options.IndexOptions{
					Unique: &isUnique,
				},
			},
		},
	)
	return &urlStore{
		collection: collection,
	}, nil
}

func (u *urlStore) CheckIfExists(ctx context.Context, url string) (bool, error) {
	opts := options.Count()
	opts.SetLimit(1)
	count, err := u.collection.CountDocuments(ctx, bson.M{"original": url}, opts)
	if err != nil {
		return true, err
	}
	if count == 0 {
		return false, nil
	}
	return true, nil
}

func (u *urlStore) Create(ctx context.Context, urlInfo service.UrlInfo) error {
	_, err := u.collection.InsertOne(ctx, urlInfo)
	return err
}

func (u *urlStore) Get(ctx context.Context, url string) (string, error) {
	opts := options.FindOne()
	opts.SetProjection(bson.M{"original": 1, "_id": 0})

	var res bson.M
	if err := u.collection.FindOne(ctx, bson.M{"shortened": url}, opts).
		Decode(&res); err != nil {
		return "", err
	}
	str, isString := res["original"].(string)
	if !isString {
		return "", errors.New("invalid url")
	}
	return str, nil
}

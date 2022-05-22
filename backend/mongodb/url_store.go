package mongodb

import (
	"context"
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
				Keys: bson.D{ {"original", 1}, {"shortened", 1}},
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

func (u *urlStore) Get(ctx context.Context, url string) (service.UrlStruct, error) {
	opts := options.FindOne()
	opts.SetProjection(bson.M{"shortened": 1})

	var res service.UrlStruct
	err := u.collection.FindOne(ctx, bson.M{"orginal": url}, opts).Decode(&res)
	return res, err
}

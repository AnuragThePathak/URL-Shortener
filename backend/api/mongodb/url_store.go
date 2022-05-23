package mongodb

import (
	"context"
	"errors"

	"github.com/AnuragThePathak/url-shortener/backend/common"
	"github.com/AnuragThePathak/url-shortener/backend/api"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type urlStore struct {
	collection *mongo.Collection
}

func NewUrlStore(database *mongo.Database) (api.UrlStore, error) {
	context, cancel :=
		context.WithTimeout(context.Background(), common.CreateIndexTimeout)
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

func (u *urlStore) Create(ctx context.Context, urlInfo api.UrlInfo) error {
	_, err := u.collection.InsertOne(ctx, urlInfo)
	return err
}

func (u *urlStore) Get(ctx context.Context, url, urlType string) (string, error) {
	var oppositeType string
	if urlType == common.ShortenedType {
		oppositeType = common.OrginalType
	} else {
		oppositeType = common.ShortenedType
	}
	opts := options.FindOne()
	opts.SetProjection(bson.M{oppositeType: 1, "_id": 0})

	var res bson.M
	if err := u.collection.FindOne(ctx, bson.M{urlType: url}, opts).
		Decode(&res); err != nil {
		return "", err
	}
	str, isString := res[oppositeType].(string)
	if !isString {
		return "", errors.New("invalid url")
	}
	return str, nil
}

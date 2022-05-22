package services

import (
	"context"
	"time"

	"github.com/osamingo/indigo"
)

type UrlStruct struct {
	Url string `json:"url"`
}

type UrlInfo struct {
	Original     string `json:"orginal"`
	Shortened    string `json:"shotened"`
	CreationTime int64  `json:"creation_time"`
}

type UrlService interface {
	Generate(context.Context, UrlStruct) (UrlStruct, error)

	Get(context.Context, UrlStruct) (UrlStruct, error)

	Delete(context.Context, UrlStruct) error
}

type urlService struct {
	UrlStore UrlStore
	IG       *indigo.Generator
}

func (u *urlService) Generate(
	ctx context.Context, urlStruct UrlStruct) (UrlStruct, error) {
	exists, err := u.UrlStore.CheckIfExists(ctx, urlStruct.Url)
	if err != nil {
		return UrlStruct{}, err
	}
	if exists {
		return u.Get(ctx, urlStruct)
	}

	id, err := u.IG.NextID()
	if err != nil {
		return UrlStruct{}, err
	}
	if err := u.UrlStore.Create(ctx,
		UrlInfo{
			Original:     urlStruct.Url,
			Shortened:    id,
			CreationTime: time.Now().Unix(),
		}); err != nil {
		return UrlStruct{}, err
	}
	return UrlStruct{Url: id}, nil
}

func (u *urlService) Get(
	ctx context.Context, urlStruct UrlStruct) (UrlStruct, error) {
	originalUrl, err := u.UrlStore.Get(ctx, urlStruct.Url)
	if err != nil {
		return UrlStruct{}, err
	}
	return UrlStruct{Url: originalUrl}, nil
}

type UrlStore interface {
	CheckIfExists(context.Context, string) (bool, error)

	Create(context.Context, UrlInfo) error

	Get(context.Context, string) (string, error)

	Delete(context.Context) error
}

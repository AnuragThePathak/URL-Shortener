package services

import (
	"context"

	"github.com/osamingo/indigo"
)

type UrlStruct struct {
	Url string `json:"url"`
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
	if err := u.UrlStore.Create(ctx, id); err != nil {
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

	Create(context.Context, string) error

	Get(context.Context, string) (string, error)

	Delete(context.Context) error
}

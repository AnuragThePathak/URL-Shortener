package services

import "context"

type UrlStruct struct {
	Url string `json:"url"`
}

type UrlService interface {
	Generate(context.Context, UrlStruct) (UrlStruct, error)

	Get(context.Context, UrlStruct) (UrlStruct, error)

	Delete(context.Context, UrlStruct) error
}

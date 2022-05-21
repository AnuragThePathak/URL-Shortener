package services

import "context"

type UrlStruct struct {
	Url string `json:"url"`
}

type UrlService interface {
	Generate(context.Context, string) (UrlStruct, error)

	Get(context.Context, string) (UrlStruct, error)

	Delete(context.Context, string) error
}

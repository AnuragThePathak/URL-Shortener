package api

import (
	"context"
	"fmt"
	"net"
	"net/url"
	"time"

	"github.com/AnuragThePathak/url-shortener/backend/common"
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
}

type urlService struct {
	urlStore UrlStore
	ig       *indigo.Generator
}

func NewUrlService(urlStore UrlStore, ig *indigo.Generator) UrlService {
	return &urlService{
		urlStore: urlStore,
		ig:       ig,
	}
}

func (u *urlService) Generate(
	ctx context.Context, urlStruct UrlStruct) (UrlStruct, error) {
	if urlObject, err := url.ParseRequestURI(urlStruct.Url); err != nil {
		return UrlStruct{}, fmt.Errorf("error parsing url: %w", err)
	} else {
		if urlObject.Scheme != "http" && urlObject.Scheme != "https" {
			return UrlStruct{}, fmt.Errorf("url scheme must be http or https %w",
				&url.Error{URL: urlStruct.Url})
		}
		if _, err := net.LookupIP(urlObject.Host); err != nil {
			return UrlStruct{}, fmt.Errorf("error looking up host: %w",
				&url.Error{URL: urlStruct.Url})
		}
	}

	exists, err := u.urlStore.CheckIfExists(ctx, urlStruct.Url)
	if err != nil {
		return UrlStruct{}, fmt.Errorf("error checking if url exists: %w", err)
	}
	if exists {
		res, err := u.urlStore.Get(ctx, urlStruct.Url, common.OrginalType)
		if err != nil {
			return UrlStruct{}, fmt.Errorf("error getting url: %w", err)
		}
		return UrlStruct{Url: res}, nil
	}

	id, err := u.ig.NextID()
	if err != nil {
		return UrlStruct{}, fmt.Errorf("error generating id: %w", err)
	}
	if err := u.urlStore.Create(ctx,
		UrlInfo{
			Original:     urlStruct.Url,
			Shortened:    id,
			CreationTime: time.Now().Unix(),
		}); err != nil {
		return UrlStruct{}, fmt.Errorf("error creating url: %w", err)
	}
	return UrlStruct{Url: id}, nil
}

func (u *urlService) Get(
	ctx context.Context, urlStruct UrlStruct,
) (UrlStruct, error) {
	res, err := u.urlStore.Get(ctx, urlStruct.Url, common.ShortenedType)
	if err != nil {
		return UrlStruct{}, fmt.Errorf("error getting url: %w", err)
	}
	return UrlStruct{Url: res}, nil
}

type UrlStore interface {
	CheckIfExists(context.Context, string) (bool, error)

	Create(context.Context, UrlInfo) error

	Get(context.Context, string, string) (string, error)
}

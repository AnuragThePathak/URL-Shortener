package api

import (
	"context"
	"github.com/AnuragThePathak/url-shortener/backend/common"
	"github.com/osamingo/indigo"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestUrlService_Get(t *testing.T) {
	var (
		complete  = "http://127.0.0.1/test"
		shortened = "home"
	)
	us := NewUrlService(mockUrlStore{
		get: func(ctx context.Context, url string, utype string) (string, error) {
			require.Equal(t, common.ShortenedType, utype)
			require.Equal(t, url, shortened)
			return complete, nil
		},
	}, nil)

	urlstr, err := us.Get(context.TODO(), UrlStruct{Url: shortened})

	require.NoError(t, err)
	require.Equal(t, urlstr.Url, complete)
}

func TestUrlService_Generate(t *testing.T) {
	var (
		complete = "http://127.0.0.1/test"
	)
	gen := indigo.New(nil, indigo.StartTime(time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)))

	us := NewUrlService(mockUrlStore{
		checkIfExists: func(ctx context.Context, url string) (bool, error) {
			return false, nil
		},
		create: func(ctx context.Context, info UrlInfo) error {
			require.NotEmpty(t, info.Shortened)
			require.NotZero(t, info.CreationTime)
			require.Equal(t, complete, info.Original)
			return nil
		},
	}, gen)

	urlstr, err := us.Generate(context.TODO(), UrlStruct{
		Url: complete,
	})

	require.NoError(t, err)
	require.NotEqual(t, urlstr.Url, complete)
}

type mockUrlStore struct {
	get           func(ctx context.Context, s string, s2 string) (string, error)
	create        func(ctx context.Context, info UrlInfo) error
	checkIfExists func(ctx context.Context, s string) (bool, error)
}

func (m mockUrlStore) CheckIfExists(ctx context.Context, s string) (bool, error) {
	return m.checkIfExists(ctx, s)
}

func (m mockUrlStore) Create(ctx context.Context, info UrlInfo) error {
	return m.create(ctx, info)
}

func (m mockUrlStore) Get(ctx context.Context, url string, utype string) (string, error) {
	return m.get(ctx, url, utype)
}

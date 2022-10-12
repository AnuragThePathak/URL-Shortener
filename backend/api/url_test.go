package api

import (
	"context"
	"github.com/osamingo/indigo"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestUrlService_Get(t *testing.T) {
	var (
		complete  = "http://127.0.0.1/test"
		shortened = "home"
	)

	mockStore := &MockUrlStore{}
	mockStore.Mock.Test(t)

	mockStore.On("Get", mock.Anything, mock.Anything, mock.Anything).
		Return(complete, nil)

	us := NewUrlService(mockStore, nil)

	urlstr, err := us.Get(context.TODO(), UrlStruct{Url: shortened})

	require.NoError(t, err)
	require.Equal(t, urlstr.Url, complete)
}

func TestUrlService_Generate(t *testing.T) {
	var (
		complete = "http://127.0.0.1/test"
	)
	gen := indigo.New(nil, indigo.StartTime(time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)))

	mockStore := &MockUrlStore{}
	mockStore.Mock.Test(t)

	mockStore.On("CheckIfExists", mock.Anything, mock.Anything).
		Return(false, nil)

	mockStore.On("Create", mock.Anything, mock.Anything).
		Return(nil)

	us := NewUrlService(mockStore, gen)

	urlstr, err := us.Generate(context.TODO(), UrlStruct{
		Url: complete,
	})

	require.NoError(t, err)
	require.NotEqual(t, urlstr.Url, complete)
}

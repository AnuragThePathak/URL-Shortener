package endpoints

import (
	"net/http"

	"github.com/AnuragThePathak/url-shortener/backend/api"
	"github.com/AnuragThePathak/url-shortener/backend/server"
	"github.com/go-chi/chi/v5"
)

type URLEndpoints struct {
	Service api.UrlService
}

func (u *URLEndpoints) Register(r chi.Router) {
	r.Post("/", u.Generate)
	r.Get("/{id}", u.Get)
}

func (u *URLEndpoints) Generate(w http.ResponseWriter, r *http.Request) {
	url := api.UrlStruct{}
	server.ServeRequest(
		server.InternalRequest{
			W:          w,
			R:          r,
			ReqBodyObj: &url,
			EndpointLogic: func() (interface{}, error) {
				return u.Service.Generate(r.Context(), url)
			},
			SuccessCode: http.StatusCreated,
		},
	)
}

func (u *URLEndpoints) Get(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	url := api.UrlStruct{
		Url: id,
	}
	server.ServeRequest(
		server.InternalRequest{
			W: w,
			R: r,
			EndpointLogic: func() (interface{}, error) {
				return u.Service.Get(r.Context(), url)
			},
			SuccessCode: http.StatusTemporaryRedirect,
		},
	)
}

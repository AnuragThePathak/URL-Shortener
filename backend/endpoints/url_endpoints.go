package endpoints

import (
	"net/http"

	"github.com/AnuragThePathak/url-shortener/backend/server"
	"github.com/AnuragThePathak/url-shortener/backend/services"
	"github.com/go-chi/chi/v5"
)

type UrlEndpoints struct {
	Service services.UrlService
}

func (u *UrlEndpoints) Register(r chi.Router) {
	r.Post("/", u.Generate)
	r.Get("/{id}", u.Get)
}

func (u *UrlEndpoints) Generate(w http.ResponseWriter, r *http.Request) {
	url := services.UrlStruct{}
	server.ServeRequest(
		server.InboundRequest{
			W:          w,
			R:          r,
			ReqBodyObj: &url,
			EndpointLogic: func() (interface{}, error) {
				return u.Service.Generate(r.Context(), url)
			},
		},
	)
}

func (u *UrlEndpoints) Get(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	url := services.UrlStruct{
		Url: id,
	}
	server.ServeRequest(
		server.InboundRequest{
			W: w,
			R: r,
			ReqBodyObj: &url,
			EndpointLogic: func() (interface{}, error) {
				return u.Service.Get(r.Context(), url)
			},
		},
	)
}

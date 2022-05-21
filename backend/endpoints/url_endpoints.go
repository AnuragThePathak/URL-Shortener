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
	r.Post("/url", u.Generate)
}

func (u *UrlEndpoints) Generate(w http.ResponseWriter, r *http.Request) {
	url := services.UrlStruct{}
	server.ServeRequest(
		server.InboundRequest{
			W:          w,
			R:          r,
			ReqBodyObj: &url,
			EndpointLogic: func() (interface{}, error) {
				return u.Service.Generate(r.Context(), url.Url)
			},
		},
	)
}

package server

import "github.com/go-chi/chi"

type Endpoints interface {
	Register(r chi.Router)
}

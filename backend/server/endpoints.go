package server

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"net/url"

	"github.com/go-chi/chi/v5"
	"go.mongodb.org/mongo-driver/mongo"
)

type Endpoints interface {
	Register(r chi.Router)
}

func ReadRequestBody(
	w http.ResponseWriter,
	r *http.Request,
	bodyObj interface{},
) bool {
	if bodyObj != nil {
		if err := json.NewDecoder(r.Body).Decode(&bodyObj); err != nil {
			WriteAPIResponse(w, http.StatusBadRequest, nil)
			return false
		}
	}
	return true
}

func ServeRequest(req InboundRequest) {
	if req.ReqBodyObj != nil {
		if !ReadRequestBody(req.W, req.R, req.ReqBodyObj) {
			return
		}
	}

	respBodyObj, err := req.EndpointLogic()
	if err != nil {
		switch {
		case errors.Is(err, context.Canceled), errors.Is(err,
			context.DeadlineExceeded):
		case errors.Is(err, mongo.ErrNoDocuments):
			WriteAPIResponse(req.W, http.StatusNotFound, err)
		default:
			if _, ok := err.(*url.Error); ok {
				WriteAPIResponse(req.W, http.StatusBadRequest, err)
			} else {
				WriteAPIResponse(req.W, http.StatusInternalServerError, err)
			}
		}
		return
	}

	WriteAPIResponse(req.W, req.SuccessCode, respBodyObj)
}

func WriteAPIResponse(
	w http.ResponseWriter,
	statusCode int,
	res interface{},
) {
	if statusCode == 0 {
		statusCode = http.StatusOK
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if res != nil {
		if err := json.NewEncoder(w).Encode(res); err != nil {
			log.Println(err)
		}
	}
}

package server

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/pkg/errors"
)

type Endpoints interface {
	Register(r chi.Router)
}

func ReadAndValidateRequestBody(
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
		if !ReadAndValidateRequestBody(req.W, req.R, req.ReqBodyObj) {
			return
		}
	}

	respBodyObj, err := req.EndpointLogic()
	if err != nil {
		switch {
		case errors.Is(err, context.Canceled), errors.Is(err,
			context.DeadlineExceeded):
			return
		default:
			WriteAPIResponse(req.W, http.StatusInternalServerError, err)
			return
		}
	}

	WriteAPIResponse(req.W, req.SuccessCode, respBodyObj)
}

func WriteAPIResponse(
	w http.ResponseWriter,
	statusCode int,
	response interface{},
) {
	if statusCode == 0 {
		statusCode = http.StatusOK
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if response != nil {
		if err := json.NewEncoder(w).Encode(response); err != nil {
			log.Println(errors.Wrap(err, "error marshaling response"))
		}
	}
}

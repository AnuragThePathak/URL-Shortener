package server

import (
	"net/http"
)

type InternalRequest struct {
	W             http.ResponseWriter
	R             *http.Request
	ReqBodyObj    interface{}
	EndpointLogic func() (interface{}, error)
	SuccessCode   int
}

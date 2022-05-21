package main

import (
	"log"

	"github.com/AnuragThePathak/url-shortener/backend/server"
)

func main() {

	var apiserver server.Server
	{
		apiserverConfig, err := serverConfig()
		if err != nil {
			log.Fatal(err)
		}
		apiserver = server.NewServer([]server.Endpoints{
		}, &apiserverConfig)
	}

	apiserver.ListenAndServe()
}

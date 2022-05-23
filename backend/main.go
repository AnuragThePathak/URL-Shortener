package main

import (
	"log"
	"time"

	"github.com/AnuragThePathak/my-go-packages/signals"
	"github.com/AnuragThePathak/url-shortener/backend/endpoints"
	"github.com/AnuragThePathak/url-shortener/backend/api/mongodb"
	"github.com/AnuragThePathak/url-shortener/backend/server"
	"github.com/AnuragThePathak/url-shortener/backend/api"
	"github.com/osamingo/indigo"
	"go.mongodb.org/mongo-driver/mongo"
)

func main() {
	ctx := signals.Context()
	var err error

	var database *mongo.Database
	{
		if database, err = databaseConnection(ctx); err != nil {
			log.Fatal(err)
		}
	}

	var urlStore api.UrlStore
	{
		if urlStore, err = mongodb.NewUrlStore(database); err != nil {
			log.Fatal(err)
		}
	}

	var indigoGenerator *indigo.Generator
	{
		t := time.Unix(time.Now().Unix(), 0)
		indigoGenerator = indigo.New(nil, indigo.StartTime(t))
	}

	urlService := api.NewUrlService(urlStore, indigoGenerator)

	var apiserver server.Server
	{
		apiserverConfig, err := serverConfig()
		if err != nil {
			log.Fatal(err)
		}
		apiserver = server.NewServer([]server.Endpoints{
			&endpoints.UrlEndpoints{
				Service: urlService,
			},
		}, &apiserverConfig)
	}

	apiserver.ListenAndServe()
	<-ctx.Done()
}

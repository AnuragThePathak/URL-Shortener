package main

import (
	"log"
	"time"

	"github.com/AnuragThePathak/my-go-packages/signals"
	"github.com/AnuragThePathak/url-shortener/backend/api"
	"github.com/AnuragThePathak/url-shortener/backend/api/mongodb"
	"github.com/AnuragThePathak/url-shortener/backend/endpoints"
	"github.com/AnuragThePathak/url-shortener/backend/server"
	"github.com/osamingo/indigo"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

func main() {
	ctx := signals.Context()
	var err error

	var logger *zap.Logger
	{
		config := zapConfig()
		if logger, err = config.Build(); err != nil {
			log.Fatal(err)
		}
	}

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
		apiserverConfig, err := serverConfig(logger)
		if err != nil {
			log.Fatal(err)
		}
		apiserver = server.NewServer([]server.Endpoints{
			&endpoints.URLEndpoints{
				Service: urlService,
				Logger: logger,
			},
		}, &apiserverConfig, logger)
	}

	apiserver.ListenAndServe()
	<-ctx.Done()
}

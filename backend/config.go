package main

import (
	"context"
	"log"
	"time"

	"github.com/AnuragThePathak/url-shortener/backend/server"
	"github.com/brigadecore/brigade-foundations/os"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func databaseConnection(ctx context.Context) (*mongo.Database, error) {
	dbConnectionStr, err := os.GetRequiredEnvVar("DB_URL")
	if err != nil {
		return nil, err
	}
	dbName, err := os.GetRequiredEnvVar("DB_NAME")
	if err != nil {
		return nil, err
	}
	dbClientOpts := options.Client().ApplyURI(dbConnectionStr)

	connectCtx, connectCancel := context.WithTimeout(ctx, 10*time.Second)
	defer connectCancel()
	client, err := mongo.Connect(connectCtx, dbClientOpts)
	if err != nil {
		return nil, err
	}
	return client.Database(dbName), nil
}

func serverConfig() (server.ServerConfig, error) {
	config := server.ServerConfig{}
	var err error

	config.Port, err = os.GetIntFromEnvVar("API_SERVER_PORT", 8080)
	if err != nil {
		return config, err
	}
	log.Println("API_SERVER_PORT: ", config.Port)
	config.TLSEnabled, err = os.GetBoolFromEnvVar("TLS_ENABLED", false)
	if err != nil {
		return config, err
	}
	log.Println("TLS_ENABLED: ", config.TLSEnabled)
	if config.TLSEnabled {
		config.TLSCertPath, err = os.GetRequiredEnvVar("TLS_CERT_PATH")
		if err != nil {
			return config, err
		}
		log.Println("TLS_CERT_PATH: ", config.TLSCertPath)
		config.TLSKeyPath, err = os.GetRequiredEnvVar("TLS_KEY_PATH")
		if err != nil {
			return config, err
		}
		log.Println("TLS_KEY_PATH: ", config.TLSKeyPath)
	}
	return config, nil
}

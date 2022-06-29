package main

import (
	"context"
	"log"
	"time"

	"github.com/AnuragThePathak/my-go-packages/os"
	"github.com/AnuragThePathak/url-shortener/backend/server"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
)

func databaseConnection(ctx context.Context) (*mongo.Database, error) {
	dbUrl, err := os.GetEnv("DB_URL")
	if err != nil {
		return nil, err
	}
	dbName, err := os.GetEnv("DB_NAME")
	if err != nil {
		return nil, err
	}
	dbClientOpts := options.Client().ApplyURI(dbUrl)

	connectCtx, connectCancel := context.WithTimeout(ctx, 10*time.Second)
	defer connectCancel()
	client, err := mongo.Connect(connectCtx, dbClientOpts)
	if err != nil {
		return nil, err
	}
	return client.Database(dbName), nil
}

func serverConfig(logger *zap.Logger) (server.ServerConfig, error) {
	config := server.ServerConfig{}
	var err error

	config.Port, err = os.GetEnvAsInt("PORT", 8080)
	if err != nil {
		return config, err
	}
	logger.Debug("PORT:", zap.Int("port", config.Port))
	config.TLSEnabled, err = os.GetEnvAsBool("TLS_ENABLED", false)
	if err != nil {
		return config, err
	}
	logger.Debug("TLS_ENABLED:", zap.Bool("tlsEnabled", config.TLSEnabled))
	if config.TLSEnabled {
		config.TLSCertPath, err = os.GetEnv("TLS_CERT_PATH")
		if err != nil {
			return config, err
		}
		logger.Debug("TLS_CERT_PATH:", zap.String("tlsCertPath",
			config.TLSCertPath))
		config.TLSKeyPath, err = os.GetEnv("TLS_KEY_PATH")
		if err != nil {
			return config, err
		}
		logger.Debug("TLS_KEY_PATH:", zap.String("tlsKeyPath", config.TLSKeyPath))
	}
	return config, nil
}

func zapConfig() zap.Config {
	env, _ := os.GetEnv("ENV")
	log.Println("ENV:", env)
	if env == "production" {
		return zap.NewProductionConfig()
	}
	return zap.NewDevelopmentConfig()
}

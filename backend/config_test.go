package main

import (
	"context"
	"os"
	"testing"

	"github.com/AnuragThePathak/url-shortener/backend/server"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

func Test_databaseConnection(t *testing.T) {
	tests := []struct {
		setup      func()
		assertions func(*mongo.Database, error)
		name       string
	}{
		{
			name:  "DB_URL not set",
			setup: func() {},
			assertions: func(_ *mongo.Database, err error) {
				require.Error(t, err)
				require.Contains(t, err.Error(), "is not set")
				require.Contains(t, err.Error(), "DB_URL")
			},
		},
		{
			name: "DB_NAME not set",
			setup: func() {
				os.Setenv("DB_URL", "mongodb://xyz")
			},
			assertions: func(_ *mongo.Database, err error) {
				require.Error(t, err)
				require.Contains(t, err.Error(), "is not set")
				require.Contains(t, err.Error(), "DB_NAME")
			},
		},
		{
			name: "success",
			setup: func() {
				os.Setenv("DB_URL", "mongodb://xyz")
				os.Setenv("DB_NAME", "new")
			},
			assertions: func(d *mongo.Database, err error) {
				require.NoError(t, err)
				require.NotNil(t, d)
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(_ *testing.T) {
			test.setup()
			db, err := databaseConnection(context.Background())
			test.assertions(db, err)
		})
	}
}

func Test_serverConfig(t *testing.T) {
	tests := []struct {
		setup      func()
		assertions func(server.ServerConfig, error)
		name       string
	}{
		{
			name: "PORT is not int",
			setup: func() {
				os.Setenv("PORT", "foo")
			},
			assertions: func(_ server.ServerConfig, err error) {
				require.Error(t, err)
				require.Contains(t, err.Error(), "can't be parsed as an integer")
				require.Contains(t, err.Error(), "PORT")
			},
		},
		{
			name: "TLS_ENABLED is not bool",
			setup: func() {
				os.Setenv("PORT", "8080")
				os.Setenv("TLS_ENABLED", "foo")
			},
			assertions: func(_ server.ServerConfig, err error) {
				require.Error(t, err)
				require.Contains(t, err.Error(), "can't be parsed as a boolean")
				require.Contains(t, err.Error(), "TLS_ENABLED")
			},
		},
		{
			name: "TLS_ENABLED is false and success",
			setup: func() {
				os.Setenv("TLS_ENABLED", "false")
			},
			assertions: func(sc server.ServerConfig, err error) {
				require.NoError(t, err)
				require.Equal(t,
					server.ServerConfig{
						Port:       8080,
						TLSEnabled: false,
					},
					sc)
			},
		},
		{
			name: "TLS_ENABLED is true and TLS_CERT_PATH not set",
			setup: func() {
				os.Setenv("TLS_ENABLED", "true")
			},
			assertions: func(_ server.ServerConfig, err error) {
				require.Error(t, err)
				require.Contains(t, err.Error(), "is not set")
				require.Contains(t, err.Error(), "TLS_CERT_PATH")
			},
		},
		{
			name: "TLS_ENABLED is true and TLS_KEY_PATH not set",
			setup: func() {
				os.Setenv("TLS_CERT_PATH", "/x/cert")
			},
			assertions: func(_ server.ServerConfig, err error) {
				require.Error(t, err)
				require.Contains(t, err.Error(), "is not set")
				require.Contains(t, err.Error(), "TLS_KEY_PATH")
			},
		},
		{
			name: "TLS_ENABLED is true and success",
			setup: func() {
				os.Setenv("TLS_KEY_PATH", "/x/key")
			},
			assertions: func(sc server.ServerConfig, err error) {
				require.NoError(t, err)
				require.Equal(t,
					server.ServerConfig{
						Port:        8080,
						TLSEnabled:  true,
						TLSCertPath: "/x/cert",
						TLSKeyPath:  "/x/key",
					}, sc)
			},
		},
	}

	logger, err := zap.NewDevelopment()
	if err != nil {
		t.FailNow()
	}

	for _, test := range tests {
		t.Run(test.name, func(_ *testing.T) {
			test.setup()
			config, err := serverConfig(logger)
			test.assertions(config, err)
		})
	}
}

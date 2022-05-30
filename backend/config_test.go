package main

import (
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/mongo"
)

func Test_databaseConnection(t *testing.T) {
	tests := []struct {
		setupDB    func()
		assertions func(*mongo.Database, error)
		name       string
	}{
		{
			name:    "DB_URL not set",
			setupDB: func() {},
			assertions: func(_ *mongo.Database, err error) {
				require.Error(t, err)
				require.Contains(t, err.Error(), "value not found for")
				require.Contains(t, err.Error(), "DB_URL")
			},
		},
		{
			name: "DB_NAME not set",
			setupDB: func() {
				os.Setenv("DB_URL", "mongodb://xyz")
			},
			assertions: func(_ *mongo.Database, err error) {
				require.Error(t, err)
				require.Contains(t, err.Error(), "value not found for")
				require.Contains(t, err.Error(), "DB_NAME")
			},
		},
		{
			name: "success",
			setupDB: func() {
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
			test.setupDB()
			db, err := databaseConnection(context.Background())
			test.assertions(db, err)
		})
	}
}

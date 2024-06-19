package mongodb

import (
	"log/slog"

	"go.mongodb.org/mongo-driver/mongo"
)

type FindrDB struct {
	DB *mongo.Client
	Logger *slog.Logger
}

// NewLinkerDB returns a new constructor instance of the CompliantDB struct.
func NewFindrDB(db *mongo.Client, logger *slog.Logger) FindrStore {
	return &FindrDB{
		DB: db,
		Logger: logger,
	}
}

// LinkerData returns a reference to a collection in the nhia database.
func FindrData(client *mongo.Client, collectionName string) *mongo.Collection {
	return client.Database("Findr").Collection(collectionName)
}

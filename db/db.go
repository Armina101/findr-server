package db

import (
	"context"
	"log/slog"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)


func SetConnection(uri string, logger *slog.Logger) (*mongo.Client, error) {
	ctx, cancelCtx := context.WithTimeout(context.Background(), 100*time.Second)

	defer cancelCtx()

	serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPIOptions))
	if err != nil {
		logger.Error("error while connecting to the database")
		panic(err)
	}
	err = client.Ping(ctx, nil)
	if err != nil {
		logger.Error("cannot ping the database")
		panic(err)
	}
	return client, nil
}

func OpenConnection(uri string, logger *slog.Logger) *mongo.Client {
	// uri := "Todo: Add MongoDB URI here"
	connectCount := 0

	logger.Info("........... Setting Connection to MongoDB ...........")

	for {

		client, err := SetConnection(uri, logger)

		if err != nil {
			logger.Info(".......... MongoDB not ready for connection ..........")
			connectCount++
		} else {
			logger.Info(".......... MongoDB client Connected  ..........")
			return client
		}

		if connectCount >= 5 {
			logger.Info(err.Error())
			return nil

		}
		logger.Info(".......... MongoDB client trying to reconnect ..........")
		time.Sleep(10 * time.Second)
		continue
	}
}

package mongodb

import (
	"context"
	"log/slog"
	"time"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func ConnectMongoDb(url string) (*mongo.Client, error) {
	clientOptions := options.Client().ApplyURI(url)

	client, err := mongo.Connect(clientOptions)

	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err = client.Ping(ctx, nil); err != nil {
		cleanupCtx, cleanupCancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cleanupCancel()
		client.Disconnect(cleanupCtx)
		return nil, err
	}

	slog.Info("MongoClient connected")

	return client, nil
}

package mongodb

import (
	"context"
	"log/slog"

	"github.com/Xav147/flashcards_backend/internal/flashcards"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type MongodbRepo struct {
	collection *mongo.Collection
}

func CreateNewMongoDbRepo(collection *mongo.Collection) *MongodbRepo {
	return &MongodbRepo{
		collection: collection,
	}
}

func (repo *MongodbRepo) ListDecks(ctx context.Context) ([]flashcards.Deck, error) {

	results := make([]flashcards.Deck, 0)
	cur, err := repo.collection.Find(ctx, bson.D{})
	if err != nil {
		return results, err
	}
	defer cur.Close(ctx)
	for cur.Next(ctx) {
		var elem flashcards.Deck
		err := cur.Decode(&elem)
		if err != nil {
			return results, err
		}
		results = append(results, elem)
	}

	if err := cur.Err(); err != nil {
		return results, err
	}

	slog.Info("Found multiple decks")

	return results, nil
}

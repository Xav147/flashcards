package flashcards

import (
	"context"
)

type FlashcardsRepository interface {
	ListDecks(ctx context.Context) ([]Deck, error)
	CreateDeck(ctx context.Context, deck Deck) error
}

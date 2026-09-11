package flashcards

import (
	"context"
)

type Service interface {
	ListDecks(ctx context.Context) ([]Deck, error)
	CreateDeck(ctx context.Context, deck Deck) error
}

type svc struct {
	// repository
	repository FlashcardsRepository
}

func NewService(repo FlashcardsRepository) Service {
	return &svc{
		repository: repo,
	}
}

func (s *svc) ListDecks(ctx context.Context) ([]Deck, error) {
	decks, err := s.repository.ListDecks(ctx)
	if err != nil {
		return nil, err
	}
	return decks, nil
}

func (s *svc) CreateDeck(ctx context.Context, deck Deck) error {
	err := s.repository.CreateDeck(ctx, deck)
	if err != nil {
		return err
	}
	return nil
}

package flashcards

import (
	"context"
	"errors"
	"reflect"
	"testing"
)

type stubRepository struct {
	decks []Deck
	err   error
}

func (r stubRepository) ListDecks(context.Context) ([]Deck, error) {
	return r.decks, r.err
}

func TestServiceListDecks(t *testing.T) {
	want := []Deck{{Name: "Go", Size: 10}}
	service := NewService(stubRepository{decks: want})

	got, err := service.ListDecks(context.Background())
	if err != nil {
		t.Fatalf("ListDecks returned an unexpected error: %v", err)
	}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("ListDecks returned %#v, want %#v", got, want)
	}
}

func TestServiceListDecksReturnsRepositoryError(t *testing.T) {
	wantErr := errors.New("repository unavailable")
	service := NewService(stubRepository{err: wantErr})

	_, err := service.ListDecks(context.Background())
	if !errors.Is(err, wantErr) {
		t.Fatalf("ListDecks returned error %v, want %v", err, wantErr)
	}
}

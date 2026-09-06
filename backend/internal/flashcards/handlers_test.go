package flashcards

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

type stubService struct {
	decks []Deck
	err   error
}

func (s stubService) ListDecks(context.Context) ([]Deck, error) {
	return s.decks, s.err
}

func TestHandlerListDecks(t *testing.T) {
	want := []Deck{{Name: "MongoDB", Size: 5}}
	handler := NewHandler(stubService{decks: want})
	request := httptest.NewRequest(http.MethodGet, "/decks", nil)
	response := httptest.NewRecorder()

	handler.ListDecks(response, request)

	if response.Code != http.StatusOK {
		t.Fatalf("ListDecks returned status %d, want %d", response.Code, http.StatusOK)
	}
	if got := response.Header().Get("Content-Type"); got != "application/json" {
		t.Fatalf("ListDecks returned Content-Type %q, want application/json", got)
	}

	var got []Deck
	if err := json.NewDecoder(response.Body).Decode(&got); err != nil {
		t.Fatalf("decode response: %v", err)
	}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("ListDecks returned %#v, want %#v", got, want)
	}
}

func TestHandlerListDecksReturnsInternalServerError(t *testing.T) {
	handler := NewHandler(stubService{err: errors.New("repository unavailable")})
	request := httptest.NewRequest(http.MethodGet, "/decks", nil)
	response := httptest.NewRecorder()

	handler.ListDecks(response, request)

	if response.Code != http.StatusInternalServerError {
		t.Fatalf("ListDecks returned status %d, want %d", response.Code, http.StatusInternalServerError)
	}
}

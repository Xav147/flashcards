package flashcards

import (
	stdjson "encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/Xav147/flashcards_backend/internal/json"
)

type handler struct {
	service Service
}

func NewHandler(service Service) *handler {
	return &handler{
		service: service,
	}
}

func (h *handler) ListDecks(w http.ResponseWriter, r *http.Request) {
	// Call the service to list the products
	decks, err := h.service.ListDecks(r.Context())
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Return JSON in an HTTP response
	json.Write(w, http.StatusOK, decks)
}

func (h *handler) CreateDeck(w http.ResponseWriter, r *http.Request) {
	var deck Deck
	err := stdjson.NewDecoder(r.Body).Decode(&deck)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = h.service.CreateDeck(r.Context(), deck)
	if err != nil {
		if errors.Is(err, ErrDeckAlreadyExists) {
			http.Error(w, err.Error(), http.StatusConflict)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
	var result []int
	json.Write(w, http.StatusOK, result)
}

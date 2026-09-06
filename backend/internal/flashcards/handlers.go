package flashcards

import (
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

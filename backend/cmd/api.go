package main

import (
	"log"
	"net/http"
	"time"

	"github.com/Xav147/flashcards_backend/internal/flashcards"
	"github.com/Xav147/flashcards_backend/internal/mongodb"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type application struct {
	config config
	db     *mongo.Client
}

func (app *application) mount() http.Handler {
	r := chi.NewRouter()
	// Middleware
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Use(middleware.Timeout(60 * time.Second))

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Up!"))
	})

	collection := app.db.Database("flashcards-db").Collection("decks")
	mongodbRepo := mongodb.CreateNewMongoDbRepo(collection)
	flashcardsService := flashcards.NewService(mongodbRepo)
	flashcardsHandler := flashcards.NewHandler(flashcardsService)
	r.Get("/list_decks", flashcardsHandler.ListDecks)

	//http.ListenAndServe(":3333",r)

	return r
}

func (app *application) run(h http.Handler) error {
	srv := &http.Server{
		Addr:         app.config.addr,
		Handler:      h,
		WriteTimeout: time.Second * 30,
		ReadTimeout:  time.Second * 10,
		IdleTimeout:  time.Minute,
	}

	log.Printf("server has started at addr %s", app.config.addr)

	return srv.ListenAndServe()
}

type config struct {
	addr string
	db   dbConfig
}

type dbConfig struct {
	dsn string
}

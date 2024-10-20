package main

import (
	"fmt"
	"log"
	"net/http"
	"tic_tac_toe/handlers"
	"tic_tac_toe/models"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

func main() {
	games := models.Games{
		Games: make(map[uuid.UUID]models.Game),
	}

	router := chi.NewRouter()
	router.Route("/games", func(r chi.Router) {
		r.Post("/", handlers.CreateGameHandler(&games))
		r.Get("/{uid}", handlers.FetchGameHandler(&games))
	})

	fmt.Println("Server is starting on port 8089...")
	log.Fatal(http.ListenAndServe(":8089", router))
}

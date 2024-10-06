package main

import (
	"fmt"
	"log"
	"net/http"
	"tic_tac_toe/handlers"
	"tic_tac_toe/models"

	"github.com/google/uuid"
)

func main() {
	games := models.Games{
		Games: make(map[uuid.UUID]models.Game),
	}

	gamesHandler := handlers.NewGamesHandler(&games)
	http.HandleFunc("/games", gamesHandler)
	http.HandleFunc("/games/", gamesHandler)

	fmt.Println("Server is starting on port 8089...")
	log.Fatal(http.ListenAndServe(":8089", nil))
}

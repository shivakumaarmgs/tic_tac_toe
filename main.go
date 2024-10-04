package main

import (
	"fmt"
	"log"
	"net/http"
	"tic_tac_toe/handlers"
	"tic_tac_toe/models"
)

func main() {
	games := models.Games{
		Games: make(map[string]models.Game),
	}

	http.HandleFunc("/games", handlers.NewGamesHandler(&games))
	http.HandleFunc("/games/", handlers.NewGamesHandler(&games))

	fmt.Println("Server is starting on port 8089...")
	log.Fatal(http.ListenAndServe(":8089", nil))
}

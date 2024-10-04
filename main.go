package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/google/uuid"
)

func main() {
	http.HandleFunc("/games", gamesHandler)

	fmt.Println("Server is starting on port 8089...")
	log.Fatal(http.ListenAndServe(":8089", nil))
}

type Games struct {
	Games map[string]Game
	Count int
}

func (g *Games) addGame(game Game) {
	g.Games[game.Uuid] = game
}

type Game struct {
	Name string `json:"name"`
	Uuid string `json:"uuid"`
}

func (g *Game) generateUuid() {
	g.Uuid = uuid.New().String()
}

var games = Games{
	Games: make(map[string]Game),
}

func gamesHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	fmt.Println("Received params")
	fmt.Println(r.URL.Path)
	fmt.Println(len("/games/"))

	var game Game
	err := json.NewDecoder(r.Body).Decode(&game)
	if err != nil {
		http.Error(w, "Error while decoding request", http.StatusBadRequest)
	}

	game.generateUuid()
	games.addGame(game)

	fmt.Println(game)
	fmt.Println(games)

	fmt.Println("")

	resp, err := json.Marshal(game)
	if err != nil {
		http.Error(w, "Error while producing response", http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(resp)
}

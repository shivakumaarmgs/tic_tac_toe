package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"tic_tac_toe/models"
)

func NewGamesHandler(games *models.Games) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			createGameHandler(games, w, r)
		case http.MethodGet:
			getGameHandler(games, w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	}
}

func createGameHandler(games *models.Games, w http.ResponseWriter, r *http.Request) {
	var game models.Game
	err := json.NewDecoder(r.Body).Decode(&game)
	if err != nil {
		http.Error(w, "Error while decoding request", http.StatusBadRequest)
	}

	game.GenerateUuid()
	games.AddGame(game)

	fmt.Println(game)
	fmt.Println(games)

	fmt.Println("")

	resp, err := json.Marshal(game)
	if err != nil {
		http.Error(w, "Error while producing response", http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(resp)
}

func getGameHandler(games *models.Games, w http.ResponseWriter, r *http.Request) {
	pathSplits := strings.Split(r.URL.Path, "/")
	if len(pathSplits) != 3 {
		http.Error(w, "URL invalid", http.StatusBadRequest)
	}

	id := pathSplits[2]
	game := games.GetGame(id)

	resp, err := json.Marshal(game)
	if err != nil {
		http.Error(w, "Error while producing response", http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(resp)
}

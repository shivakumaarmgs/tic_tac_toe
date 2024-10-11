package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"tic_tac_toe/models"
	"tic_tac_toe/utils"

	"github.com/google/uuid"
)

func NewGamesHandler(games *models.Games) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			createGameHandler(games, w, r)
		case http.MethodGet:
			getGameHandler(games, w, r)
		default:
			utils.RespondWithError(w, http.StatusMethodNotAllowed, "Method not allowed")
		}
	}
}

func createGameHandler(games *models.Games, w http.ResponseWriter, r *http.Request) {
	var game models.Game
	err := json.NewDecoder(r.Body).Decode(&game)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Error while decoding request body")
		return
	}

	game.GenerateUuid()
	games.AddGame(game)
	fmt.Println("All Games")
	fmt.Println(games)
	fmt.Println(" ")

	w.Header().Set("Content-Type", "application/json")
	resp, err := json.Marshal(game)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Error while producing response")
		return
	}
	_, _ = w.Write(resp)
}

func getGameHandler(games *models.Games, w http.ResponseWriter, r *http.Request) {
	pathSplits := strings.Split(r.URL.Path, "/")
	if len(pathSplits) != 3 {
		utils.RespondWithError(w, http.StatusBadRequest, "URL invalid")
		return
	}

	id, err := uuid.Parse(pathSplits[2])
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "UUID given in the URL is not valid")
		return
	}

	game, ok := games.GetGame(id)
	if !ok {
		utils.RespondWithError(w, http.StatusNotFound, "Game not found for given uuid")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(game)
}

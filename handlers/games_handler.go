package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"tic_tac_toe/models"
	"tic_tac_toe/utils"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

func CreateGameHandler(games *models.Games) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var game models.Game
		err := json.NewDecoder(r.Body).Decode(&game)
		if err != nil {
			utils.RespondWithError(w, http.StatusBadRequest, "Error while decoding request body")
			return
		}

		game.GenerateUuid()

		validate := validator.New(validator.WithRequiredStructEnabled())
		err = validate.Struct(game)
		if err != nil {
			utils.RespondWithValidationErrors(w, err.(validator.ValidationErrors))
			return
		}

		games.AddGame(game)
		fmt.Println("All Games")
		fmt.Println(games)
		fmt.Println(" ")

		w.Header().Set("Content-Type", "application/json")
		resp, err := json.Marshal(game)
		if err != nil {
			utils.RespondWithError(
				w,
				http.StatusInternalServerError,
				"Error while producing response",
			)
			return
		}
		_, _ = w.Write(resp)
	}
}

func FetchGameHandler(games *models.Games) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		uid := chi.URLParam(r, "uid")
		id, err := uuid.Parse(uid)
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
}

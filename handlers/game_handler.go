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

func MarkBoxHanlder(games *models.Games) http.HandlerFunc {
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

		var markPos models.Mark
		err = json.NewDecoder(r.Body).Decode(&markPos)
		if err != nil {
			utils.RespondWithError(w, http.StatusBadRequest, "Error while decoding request body")
			return
		}
		validate := validator.New(validator.WithRequiredStructEnabled())
		err = validate.Struct(markPos)
		if err != nil {
			utils.RespondWithValidationErrors(w, err.(validator.ValidationErrors))
			return
		}
		fmt.Println(markPos)
		game.MarkPosition(markPos)

		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(game)
	}
}

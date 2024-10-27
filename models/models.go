package models

import (
	"github.com/google/uuid"
)

type Games struct {
	Games map[uuid.UUID]*Game
	Count int
}

func (g *Games) AddGame(game *Game) {
	g.Games[game.Uuid] = game
	g.Count++
}

func (g *Games) GetGame(uuid uuid.UUID) (game *Game, ok bool) {
	game, ok = g.Games[uuid]
	return
}

type Game struct {
	Name  string    `json:"name" validate:"required"`
	Uuid  uuid.UUID `json:"uuid" validate:"required"`
	Board [3][3]string
}

func (g *Game) GenerateUuid() {
	g.Uuid = uuid.New()
}

func (g *Game) InitializeBoard() {
	g.Board = [3][3]string{
		{"-", "-", "-"},
		{"-", "-", "-"},
		{"-", "-", "-"},
	}
}

func (g *Game) MarkPosition(mark Mark) {
	i := (mark.BoxNo - 1) / 3
	j := (mark.BoxNo - 1) % 3
	g.Board[i][j] = "X"
}

type Mark struct {
	BoxNo int    `json:"box_no" validate:"required"`
	Team  string `json:"team"   validate:"oneof=x_team o_team"`
}

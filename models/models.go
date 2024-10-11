package models

import (
	"github.com/google/uuid"
)

type Games struct {
	Games map[uuid.UUID]Game
	Count int
}

func (g *Games) AddGame(game Game) {
	g.Games[game.Uuid] = game
	g.Count++
}

func (g *Games) GetGame(uuid uuid.UUID) (game Game, ok bool) {
	game, ok = g.Games[uuid]
	return
}

type Game struct {
	Name string    `json:"name" validate:"required"`
	Uuid uuid.UUID `json:"uuid" validate:"required"`
}

func (g *Game) GenerateUuid() {
	g.Uuid = uuid.New()
}

package models

import "github.com/google/uuid"

type Games struct {
	Games map[uuid.UUID]Game
	Count int
}

func (g *Games) AddGame(game Game) {
	g.Games[game.Uuid] = game
	g.Count++
}

func (g *Games) GetGame(uuid uuid.UUID) (game Game) {
	return g.Games[uuid]
}

type Game struct {
	Name string    `json:"name"`
	Uuid uuid.UUID `json:"uuid"`
}

func (g *Game) GenerateUuid() {
	g.Uuid = uuid.New()
}

package models

import "github.com/google/uuid"

type Games struct {
	Games map[string]Game
	Count int
}

func (g *Games) AddGame(game Game) {
	g.Games[game.Uuid] = game
	g.Count++
}

func (g *Games) GetGame(uuid string) (game Game) {
	return g.Games[uuid]
}

type Game struct {
	Name string `json:"name"`
	Uuid string `json:"uuid"`
}

func (g *Game) GenerateUuid() {
	g.Uuid = uuid.New().String()
}

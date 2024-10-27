package models

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestGenerateUuid(t *testing.T) {
	game := Game{}

	game.GenerateUuid()

	assert.NotEqual(t, uuid.Nil, game.Uuid, "Expected a non-nil UUId")
}

func TestInitializeBoard(t *testing.T) {
	game := Game{}

	game.InitializeBoard()

	expectedBoard := [3][3]string{
		{"-", "-", "-"},
		{"-", "-", "-"},
		{"-", "-", "-"},
	}

	assert.Equal(
		t,
		expectedBoard,
		game.Board,
		"Expected board to be initialized with '-' in each cell",
	)
}

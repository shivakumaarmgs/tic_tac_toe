package handlers_test

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"tic_tac_toe/handlers"
	"tic_tac_toe/models"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type CreateGamesHandlerTestSuite struct {
	suite.Suite
}

func (s *CreateGamesHandlerTestSuite) TestCreateGamesHander() {
	s.Run("when the payload is valid", func() {
		games := models.Games{
			Games: make(map[uuid.UUID]*models.Game),
		}
		handler := handlers.CreateGamesHandler(&games)

		payload := map[string]any{
			"name": "TestGameName",
		}
		body, _ := json.Marshal(payload)

		req := httptest.NewRequest("POST", "/games", bytes.NewBuffer(body))
		recorder := httptest.NewRecorder()
		handler.ServeHTTP(recorder, req)

		assert.Equal(s.T(), http.StatusOK, recorder.Code)
	})

	s.Run("when the payload is invalid", func() {
		games := models.Games{
			Games: make(map[uuid.UUID]*models.Game),
		}
		handler := handlers.CreateGamesHandler(&games)

		payload := map[string]any{
			"dummy": "TestGameName",
		}
		body, _ := json.Marshal(payload)

		req := httptest.NewRequest("POST", "/games", bytes.NewBuffer(body))
		recorder := httptest.NewRecorder()
		handler.ServeHTTP(recorder, req)

		assert.Equal(s.T(), http.StatusUnprocessableEntity, recorder.Code)
	})
}

type ShowGamesHandlerTestSuite struct {
	suite.Suite
	games *models.Games
}

func (s *ShowGamesHandlerTestSuite) SetupTest() {
	s.games = &models.Games{
		Games: make(map[uuid.UUID]*models.Game),
	}
}

func (s *ShowGamesHandlerTestSuite) TestShowGamesHandler() {
	s.Run("When the uuid is valid", func() {
		gameUid := uuid.New()
		game := models.Game{
			Name: "TestGame",
			Uuid: gameUid,
		}
		s.games.AddGame(&game)

		handler := handlers.ShowGamesHandler(s.games)

		req := httptest.NewRequest("GET", "/games/"+gameUid.String(), nil)

		routeCtx := chi.NewRouteContext()
		routeCtx.URLParams.Add("uid", gameUid.String())
		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, routeCtx))

		recorder := httptest.NewRecorder()

		handler.ServeHTTP(recorder, req)

		assert.Equal(s.T(), http.StatusOK, recorder.Code)
	})

	s.Run("When the uuid is invalid", func() {
		gameUid := "invalid-string"

		handler := handlers.ShowGamesHandler(s.games)

		req := httptest.NewRequest("GET", "/games/"+gameUid, nil)

		routeCtx := chi.NewRouteContext()
		routeCtx.URLParams.Add("uid", gameUid)
		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, routeCtx))

		recorder := httptest.NewRecorder()

		handler.ServeHTTP(recorder, req)

		assert.Equal(s.T(), http.StatusBadRequest, recorder.Code)
	})
}

func TestGamesHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(CreateGamesHandlerTestSuite))
	suite.Run(t, new(ShowGamesHandlerTestSuite))
}

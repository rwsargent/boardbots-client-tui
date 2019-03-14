package api

import (
	"boardbots/quoridor"
	"github.com/google/uuid"
	"boardbots/util"
	"time"
)

type (
	BaseResponse struct {
		Error string `json:"error"`
	}

	BoardResponse struct {
		Board []TPiece `json:"board"`
		CurrentTurn quoridor.PlayerPosition `json:"currentTurn"`
	}

	NewUserResponse struct {
		BaseResponse
		UserId uuid.UUID `json:"userId"`
	}

	GameResponse struct {
		BaseResponse
		BoardResponse
		Players []TPlayer
		GameId uuid.UUID `json:gameId"`
	}

	GameRequest struct {
		GameId uuid.UUID `json:"gameId"`
	}

	GameListResponse struct {
		BaseResponse
		Games []TGame `json:"games"`
	}

	TPiece struct {
		Type rune `json:"type"`
		Position util.Position `json:"position"`
		Owner quoridor.PlayerPosition `json:"owner"`
	}

	TPlayer struct {
		Barriers int `json:"barriers"`
		PlayerName string `json:"playerName"`
		PawnPosition util.Position `json:"pawnPosition"`
	}

	TGame struct {
		Players []string `json:"players"`
		StartTime time.Time `json:"starTime"`
		CurrentTurn int `json:"currentTurn"`
		GameId uuid.UUID `json:"gameId"`
	}
)
package api

import (
	"boardbots/quoridor"
	"github.com/google/uuid"
	"boardbots/util"
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
	}

	GameRequest struct {
		GameId uuid.UUID `json:"gameId"`
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
)
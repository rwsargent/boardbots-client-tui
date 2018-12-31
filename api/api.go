package api

import "net/http"
type (
	Api interface {
		GetGame() (Game, error)
	}
)
func GetGame() (Game, error) {
	resp, err := http.Get("")
	if err != nil {

	}
}
package api

import (
	"net/http"
	"encoding/json"
	"bytes"
	"fmt"
	"log"
)

type (
	Api interface {
		NewUser() (BaseResponse, error)
		SignIn() (BaseResponse, error)
		GetGames() (GameListResponse, error)
		MakeGame() (GameResponse, error)
		JoinGame() (GameResponse, error)
	}

	Server struct {
		domain      string
		credentials Credentials
	}

	Credentials struct {
		Username string `json:"Username"`
		Password string `json:"Password"`
	}
)

const ApiPrefix = "/api/v0/"

func NewServer(domain string, credentials Credentials) Server {
	return Server {
		domain: domain,
		credentials: credentials,
	}
}

func (server Server) NewUser() (BaseResponse, error) {
	req := Credentials{
		Username: server.credentials.Username,
		Password: server.credentials.Password,
	}
	var resp BaseResponse
	if err := server.call("/newuser", req, &resp); err != nil {
		return BaseResponse{}, err
	} else {
		return resp, nil
	}
}

func (server Server) SignIn() (BaseResponse, error) {
	var resp BaseResponse
	err := server.call("/signin", server.credentials, &resp)
	return resp, err
}

func (server Server) GetGames(request GetGamesRequest) (GameListResponse, error) {
	var response GameListResponse
	err := server.call(ApiPrefix + "getgames", "",  &response)
	return response, err
}

func (server Server) MakeGame() (GameResponse, error) {
	var response GameResponse
	err := server.call(ApiPrefix + "makegame", "", &response)
	return response, err
}

func (server Server) call(path string, body interface{}, response interface{}) error {
	buffer := new(bytes.Buffer)
	json.NewEncoder(buffer).Encode(body)
	client := http.Client{}
	request, err := http.NewRequest("POST", server.domain + path, buffer)
	request.Header.Set("Content-Type", "application/json")
	if err != nil {
		return err
	}
	if len(server.credentials.Username) > 0 {
		request.SetBasicAuth(server.credentials.Username, server.credentials.Password)
	}
	rawResponse, err := client.Do(request)
	if err != nil {
		return err
	}
	log.Printf("Raw response: %v\n", rawResponse)
	log.Println("Raw Body: ", rawResponse.Body)
	decoder := json.NewDecoder(rawResponse.Body)
	err = decoder.Decode(response)
	if err != nil {
		fmt.Println(err.Error())
	}
	log.Println("Decoded: ", response)
	return nil
}
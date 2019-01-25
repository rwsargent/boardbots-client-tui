package api

import (
	"net/http"
	"encoding/json"
	"bytes"
	"fmt"
)

type (
	Api interface {
		NewUser(username, password string) (NewUserResponse, error)
		SignIn(username, password string) (BaseResponse, error)
	}

	Server struct {
		domain string
	}
)

func NewServer(domain string) Server {
	return Server {
		domain: domain,
	}
}

func (server Server) NewUser(username, password string) (NewUserResponse, error) {
	req := struct {
		username string
		password string
	}{
		username,
		password,
	}
	var newUserResponse NewUserResponse
	if err := call(server.domain + "/newuser", req, &newUserResponse); err != nil {
		return NewUserResponse{}, err
	} else {
		return newUserResponse, nil
	}
}

func (server Server) SignIn(username, password string) (BaseResponse, error) {
	req := struct {
		username string
		password string
	}{
		username,
		password,
	}
	var resp BaseResponse
	err := call(server.domain + "/signin", req, &resp)
	return resp, err
}

func call(path string, body interface{}, response interface{}) error {
	buffer := new(bytes.Buffer)
	json.NewEncoder(buffer).Encode(body)
	rawResponse, err := http.Post(path, "application/json", buffer)
	if err != nil {
		return err
	}
	decoder := json.NewDecoder(rawResponse.Body)
	err = decoder.Decode(response)
	if err != nil {
		fmt.Println(err.Error())
	}
	return nil
}
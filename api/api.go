package api

import (
	"net/http"
	"encoding/json"
	"bytes"
)

type (
	Api interface {
		NewUser(username, password string) BaseResponse
	}

	Server struct {
		domain string
	}
)

func (server Server) NewUser(username, password string) (BaseResponse, error) {
	req := struct {
		username string
		password string
	}{
		username,
		password,
	}
	var baseResponse BaseResponse
	if err := call(server.domain + "/newuser", req, &baseResponse); err != nil {
		return baseResponse, nil
	} else {
		return BaseResponse{}, err
	}
}

func call(path string, body interface{}, response interface{}) error {
	buffer := new(bytes.Buffer)
	json.NewEncoder(buffer).Encode(body)
	rawResponse, err := http.Post(path, "application/json", buffer)
	if err != nil {
		return err
	}
	decoder := json.NewDecoder(rawResponse.Body)
	decoder.Decode(response)
	return nil
}
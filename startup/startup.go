package startup

import (
	"bufio"
	"os"
	"fmt"
	"strconv"
	"boardbotclient/api"
	"github.com/google/uuid"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)
var Server api.Server

type InitInfo struct {
	Username string
	Password string
	GameId uuid.UUID
}

func StartUp(domain string) api.Server {
	server := authenticate(domain)
	chooseGame(server)
	return server
}

func chooseGame(server api.Server) api.TGame {
	var game *api.TGame = nil
	for game == nil{
		gameChoice := GetNumberedSelection([]string{"Start new game", "Join game", "Resume game"})
		//var game api.TGame
		switch gameChoice {
		case 1:
			gameResp, err := server.MakeGame()
			if err != nil {
				fmt.Println(err)
				break
			}
			if len(gameResp.Error) > 0 {
				fmt.Println(gameResp.Error)
				break
			}
			fmt.Println("Created game: ", gameResp.GameId)
			game = &api.TGame{}
		case 2:
			// join game
			server.GetGames()
		case 3:
			// resume game
		}
	}
	return api.TGame{}
}

func authenticate(domain string) api.Server {
	credentials, action := getCredentials()
	server := api.NewServer(domain, credentials)
	var resp api.BaseResponse
	var err error
	switch action {
	case 1:
		resp, err = server.NewUser()
	case 2:
		resp, err = server.SignIn()
	}
	if err != nil {
		fmt.Println("We encountered an error", err)
		panic(err.Error())
	}
	if len(resp.Error) != 0 {
		panic(resp.Error)
	}
	fmt.Println("Success!")
	return server
}

//func getGame(server api.Server) uuid.UUID {
//	gameResponse, err := server.GetGames()
//	if err != nil {
//		panic(err)
//	} else if len(gameResponse.Error) != 0 {
//		panic(gameResponse.Error)
//	}
//
//	displayGames := convertToDisplay(gameResponse.Games)
//	return uuid.Nil
//}
//
//func convertToDisplay(games []api.TGame) []string {
//	displayGames := make([]string, 0, len(games))
//	for _, game := range games {
//		game
//	}
//	fmt.Sprintf("")
//}

func getCredentials() (api.Credentials, int) {
	credentials := api.Credentials{}
	fmt.Println("Welcome to BoardBots - Quorridor!")
	choice := GetNumberedSelection([]string{"New User", "Sign In"})
	fmt.Print("Username: ")
	credentials.Username = ReadLine()
	fmt.Print("Password (this is an unsecure system): ")
	credentials.Password = ReadLine()
	return credentials, choice
}

func GetNumberedSelection(options []string) int {
	for {
		for idx, option := range options {
			fmt.Printf("%d) %s\n", idx+1, option)
		}
		input := ReadLine()
		selection, err := strconv.Atoi(strings.TrimSpace(input))
		if err != nil || selection > len(options) || selection <= 0 {
			fmt.Printf("Sorry, %s is not a valid option\n", input)
		} else {
			return selection
		}
	}
}

func ReadLine() string {
	input, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	return input
}

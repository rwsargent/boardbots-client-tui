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
var Server api.Api

type InitInfo struct {
	Username string
	Password string
	GameId uuid.UUID
}

func StartUp(server api.Api) InitInfo{
	info := InitInfo{}
	Server = server
	fmt.Println("Welcome to BoardBots - Quorridor!")
	choice := getNumberedSelection([]string{"New User", "Sign In"})
	if choice == 1 {
		username, password := newUser()
		info.Username = username
		info.Password = password
	} else if choice == 2 {
		username, password := signIn()
		info.Username = username
		info.Password = password
	}
	return info
}

func signIn() (string, string) {
	for {
		fmt.Print("Username: ")
		username := ReadLine()
		fmt.Print("Password (this is an unsecure system): ")
		password := ReadLine()

		resp, err := Server.SignIn(username, password)
		if err == nil && len(resp.Error) == 0 {
				fmt.Println("Signed in!")
				return username, password
		} else if len(resp.Error) > 0 {
			fmt.Println(resp.Error)
		} else {
			fmt.Println(err)
		}
	}
}

func newUser() (string, string) {
	for {
		fmt.Print("Username: ")
		username := ReadLine()
		fmt.Print("Password (this is an unsecure system): ")
		password := ReadLine()

		resp, err := Server.NewUser(username, password)
		if err != nil {
			panic(err)
		}
		if len(resp.Error) > 0 {
			fmt.Printf("Sorry, there was an issue: %s.\n", resp.Error)
			continue
		}
		fmt.Println("Success")
		return username, password
	}
}

func getNumberedSelection(options []string) int {
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

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	//MainPost base steam api url
	MainPost string = "https://api.steampowered.com/"
)

type mainstruct struct {
	Response ResponseStruct `json:"response"`
}

// ResponseStruct model
type ResponseStruct struct {
	Players []PlayersStruct `json:"players"`
}

// PlayersStruct model
type PlayersStruct struct {
	SteamID    string `json:"steamid"`
	ProfileURL string `json:"profileurl"`
}

func errcheck(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

//GetPlayerSummaries: Input SteamID and collect user's all infos
func GetPlayerSummaries(SteamID string, PrivyKey string) {
	steamid := "76561197960361544"
	resp, _ := http.Get("https://api.steampowered.com/ISteamUser/GetPlayerSummaries/v2/?key=" + string(PrivKey) + "&steamids=" + string(steamid))

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	textBytes := []byte(body)
	main := mainstruct{}
	err = json.Unmarshal(textBytes, &main)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(main.Response.Players)
}

//GetPlayerBans input user's SteamID and get user's ban info
func GetPlayerBans(SteamID string, PrivyKey string) {
	resp, err := http.Get(MainPost + "ISteamUser/GetPlayerBans/v1/?key=" + string(PrivyKey) + "&steamids=" + SteamID)
	errcheck(err)
	//We Read the response body on the line below.
	body, err := ioutil.ReadAll(resp.Body)
	errcheck(err)
	//Convert the body to type string
	sb := string(body)
	log.Printf(sb)
}

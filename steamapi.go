package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

var (
	steamıd string
)

const (
	//PrivKey steam api key
	PrivKey string = ""
	//MainPost base steam api url
	MainPost string = "https://api.steampowered.com/"
)

func main() {
	steamıd = "76561197960361544"
	resp, err := http.Get(MainPost + "ISteamUser/GetPlayerSummaries/v2/?key=" + string(PrivKey) + "&steamids=" + steamıd)
	if err != nil {
		log.Fatalln(err)
	}
	//We Read the response body on the line below.
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	//Convert the body to type string
	sb := string(body)
	log.Printf(sb)
}

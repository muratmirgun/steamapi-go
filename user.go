package steamapi

import (
	"io/ioutil"
	"log"
	"net/http"
)

const (
	//MainPost base steam api url
	MainPost string = "https://api.steampowered.com/"
)

func errcheck(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

//GetPlayerSummaries: Input steamıd and collect user's all infos
func GetPlayerSummaries(steamids string, PrivKey string) {
	resp, err := http.Get(MainPost + "ISteamUser/GetPlayerSummaries/v2/?key=" + string(PrivKey) + "&steamids=" + steamids)
	errcheck(err)
	//We Read the response body on the line below.
	body, err := ioutil.ReadAll(resp.Body)
	errcheck(err)
	//Convert the body to type string
	sb := string(body)
	log.Printf(sb)
}

//GetPlayerBans input user'ssteam id and get user's ban info
func GetPlayerBans(steamids string, PrivKey string) {
	resp, err := http.Get(MainPost + "ISteamUser/GetPlayerBans/v1/?key=" + string(PrivKey) + "&steamids=" + steamids)
	errcheck(err)
	//We Read the response body on the line below.
	body, err := ioutil.ReadAll(resp.Body)
	errcheck(err)
	//Convert the body to type string
	sb := string(body)
	log.Printf(sb)
}

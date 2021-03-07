package steamapi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Player struct {
	SteamId                  string `json:"steamid"`
	CommunityVisibilityState int    `json:"communityvisibilitystate"`
	ProfileState             int    `json:"profilestate"`
	PersonaName              string `json:"personaname"`
	LastLogoff               int    `json:"lastlogoff"`
	ProfileUrl               string `json:"profileurl"`
	Avatar                   string `json:"avatar"`
	AvatarMedium             string `json:"avatarmedium"`
	AvatarFull               string `json:"avatarfull"`
	PersonaState             int    `json:"personastate"`
	RealName                 string `json:"realname"`
	PrimaryClanId            int    `json:"primaryclanid"`
	TimeCreated              int    `json:"timecreated"`
	PersonaStateFlags        int    `json:"personastateflags"`
	LocCountryCode           string `json:"loccountrycode"`
	LocStateCode             string `json:"locstatecode"`
	LocCityId                int    `json:"loccityid"`
}
type Bans struct {
}
type Mainstruct struct {
	Response GetPlayerSummariesResponseBody
}

type GetPlayerSummariesResponseBody struct {
	Players []Player
	BanInfo []Bans
}

//GetPlayerSummaries: Input SteamID and collect user's all infos
func GetPlayerSummaries(SteamID string, PrivyKey string) []Player {
	resp, _ := http.Get("https://api.steampowered.com/ISteamUser/GetPlayerSummaries/v2/?key=" + string(PrivKey) + "&steamids=" + string(steamid))

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	textBytes := []byte(body)
	main := Mainstruct{}
	err = json.Unmarshal(textBytes, &main)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return main.Response.Players
}

//GetPlayerBans input user's SteamID and get user's ban info
func GetPlayerBans(SteamID string, PrivyKey string) []Bans {
	resp, err := http.Get(MainPost + "ISteamUser/GetPlayerBans/v1/?key=" + string(PrivyKey) + "&steamids=" + SteamID)
	if err != nil {
		log.Println(err)
	}
	//We Read the response body on the line below.
	body, err := ioutil.ReadAll(resp.Body)
	textBytes := []byte(body)
	main := Mainstruct{}
	err = json.Unmarshal(textBytes, &main)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return main.Response.BanInfo
}

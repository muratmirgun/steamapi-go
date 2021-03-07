package steamapi

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
	PrivKey         = ""
)

func main() {
	steamid := "76561197960361544"

	fmt.Println(main.Response.Players)

}

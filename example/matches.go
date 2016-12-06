package main

import (
	"fmt"

	"encoding/json"
	"os"

	"github.com/fivetwentysix/dota2"
)

func main() {
	dReq := dota2.NewDota2Request(os.Getenv("STEAM_API_KEY"))
	// res, err := dReq.GetLeagueListing()
	// res, err := dReq.GetLiveLeagueGames()
	res, err := dReq.GetMatchDetails(2826037294)
	if err != nil {
		fmt.Print(err)
	}

	output, _ := json.MarshalIndent(res, "", "  ")

	fmt.Println(string(output))
}

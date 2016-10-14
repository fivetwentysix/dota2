package main

import (
	"fmt"

	"github.com/jesusgalvan/dota2"
)

func main() {
	dReq := dota2.NewDota2Request("B517D45BAA8299DF3BFFFB67066A41B4")
	// res, err := dReq.GetLeagueListing()
	// res, err := dReq.GetLiveLeagueGames()
	res, err := dReq.GetMatchDetails(2589101809)
	if err != nil {
		fmt.Print(err)
	}

	fmt.Printf("%+v", res)
}

# Dota 2

This is a wrapper for the  [Dota2 API](https://wiki.teamfortress.com/wiki/WebAPI#Dota_2).

Example:

```
package main

import (
        "fmt"

        "github.com/jesusgalvan/dota2"
)

func main() {
        dReq := dota2.NewDota2Request("YOUR_API_KEY")
        // res, err := dReq.GetLeagueListing()
        // res, err := dReq.GetLiveLeagueGames()
        res, err := dReq.GetMatchDetails(2589101809)
        if err != nil {
                fmt.Print(err)
        }

        fmt.Printf("%+v", res)
}
```
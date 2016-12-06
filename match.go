package dota2

import "fmt"

// IDOTA2Match_<ID>

// GetLeagueListing returns information about
// DotaTV-supported leagues
func (d *Dota2Request) GetLeagueListing() (*Dota2Response, error) {
	d.APIEnd = "/IDOTA2Match_570/GetLeagueListing/v1/"
	res, err := d.RequestSend()
	if err != nil {
		return nil, err
	}
	return res, nil
}

// GetLiveLeagueGames returns a list of in-progress
// league matches, as well as details of that match as
// it unfolds
func (d *Dota2Request) GetLiveLeagueGames() (*Dota2Response, error) {
	d.APIEnd = "/IDOTA2Match_570/GetLiveLeagueGames/v1/"
	res, err := d.RequestSend()
	if err != nil {
		return nil, err
	}
	return res, nil
}

// GetMatchDetails return information about a
// particular match
func (d *Dota2Request) GetMatchDetails(id int) (*Dota2Response, error) {
	d.APIEnd = "/IDOTA2Match_570/GetMatchDetails/v001/"
	d.Param = fmt.Sprintf("match_id=%d", id)
	res, err := d.RequestSend()
	if err != nil {
		return nil, err
	}
	return res, nil
}

// GetMatchHistory returns a list of matches,
// filterable by various parameters
func (d *Dota2Request) GetMatchHistory() (*Dota2Response, error) {
	d.APIEnd = "/IDOTA2Match_570/GetMatchHistory/V001/"
	res, err := d.RequestSend()
	if err != nil {
		return nil, err
	}
	return res, nil
}

// GetMatchHistoryBySequenceNum returns a list of
// matches ordered by their sequence num
func (d *Dota2Request) GetMatchHistoryBySequenceNum() {

}

// GetScheduledLeagueGames returns a list of
// scheduled league games coming up
func (d *Dota2Request) GetScheduledLeagueGames() {

}

// GetTeamInfoByTeamID returns a list of
// all teams set up in-game
func (d *Dota2Request) GetTeamInfoByTeamID() {

}

// GetTournamentPlayerStats returns stats about
// a particular player within a tournament
func (d *Dota2Request) GetTournamentPlayerStats() {

}

// GetTopLiveGame
// TODO
func (d *Dota2Request) GetTopLiveGame() {

}

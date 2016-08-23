package dota2

// Dota2Response is the struct for the Steam
// response
type Dota2Response struct {
	Result Result `json:"result"`
}

type Result struct {
	Leagues []League `json:"leagues"`
	Games   []Game
}

type League struct {
	Name          string `json:"name"`
	LeagueId      int    `json:"leagueid"`
	Description   string `json:"description"`
	TournamentURL string `json:"tournament_url"`
	ItemDef       string `json:"itemdef"`
}

type Game struct {
	Players           []Player // list
	RadiantTeam       Team     `json:"radiant_team"`
	DireTeam          Team     `json:"dire_team"`
	LobbyId           int      `json:"lobby_id"`
	MatchId           int      `json:"match_id"`
	Spectators        int      `json:"spectators"`
	SeriesId          int      `json:"series_id"`
	GameNumber        int      `json:"game_number"`
	LeagueId          int      `json:"league_id"`
	TowerState        int      // 32 bit unsigned int
	StreamDelay       int      `json:"stream_delay_s"`
	RadiantSeriesWins int      `json:"radiant_series_wins"`
	DireSeriesWins    int      `json:"dire_series_wins"`
	SeriesType        int      `json:"series_type"`
	LeagueSeriesId    int      `json:"league_series_id"`
	LeagueGameId      int      `json:"league_game_id"`
	StageName         string   `json:"stage_name"`
	LeagueTier        int      `json:"league_tier"`
	Scoreboard        struct {
		Duration           float64 `json:"duration"`
		RoshanRespawnTimer int     `json:"roshan_respawn_timer"`
		Radiant            Team    `json:"radiant"`
		Dire               Team    `json:"dire"`
	} `json:"scoreboard"`
}

type Team struct {
	// IDOTA2Match_570/GetLiveLeagueGames
	TeamName string `json:"team_name"`
	TeamId   int    `json:"team_id"`
	TeamLogo int64  `json:"team_logo"`
	Complete bool   `json:"complete"`

	Score         int       `json:"score"`
	TowerState    int       `json:"tower_state"`
	BarracksState int       `json:"barracks_state"`
	Picks         []Pick    `json:"picks"`
	Bans          []Ban     `json:"picks"`
	Players       []Player  `json:"players"`
	Abilities     []Ability `json:"abilities"`
}

type Player struct {
	AccountId        int     `json:"account_id"`
	HeroId           int     `json:"hero_id"`
	PlayerSlot       int     `json:"player_slot"`
	Kills            int     `json:"kills"`
	Death            int     `json:"death"`
	Assists          int     `json:"assists"`
	LastHits         int     `json:"last_hits"`
	Denies           int     `json:"denies"`
	Gold             int     `json:"gold"`
	Level            int     `json:"level"`
	GoldPerMin       int     `json:"gold_per_min"`
	XPPerMin         int     `json:"xp_per_min"`
	UltimateState    int     `json:"ultimate_state"`
	UltimateCooldown int     `json:"ultimate_cooldown"`
	Item0            int     `json:"item0"`
	Item1            int     `json:"item1"`
	Item2            int     `json:"item2"`
	Item3            int     `json:"item3"`
	Item4            int     `json:"item4"`
	Item5            int     `json:"item5"`
	RespawnTimer     int     `json:"respawn_timer"`
	PositionX        float64 `json:"position_x"`
	PositionY        float64 `json:"position_y"`
	NetWorth         int     `json:"net_worth"`

	Team int
	Name string
}

type Pick struct {
	HeroId int `json:"hero_id"`
}

type Ban struct {
	HeroId int `json:"hero_id"`
}

type Ability struct {
	AbilityId    int `json:"ability_id"`
	AbilityLevel int `json:"ability_level"`
}

type MatchDetails struct {
	Players []Player
}

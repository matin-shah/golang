package models

import (
	"encoding/json"
	"fmt"
)

type TeamResponse struct {
	Response []Team `json:"response"`
}

type Team struct {
	ID       int           `json:"id"`
	Name     string        `json:"name"`
	Code     string        `json:"code"`
	Country  string        `json:"country"`
	Founded  int           `json:"founded"`
	National bool          `json:"national"`
	Logo     string        `json:"logo"`
	Venue    Venueresponse `json:"venue"`
}

type Venueresponse struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Address  string `json:"address"`
	City     string `json:"city"`
	Capacity int    `json:"capacity"`
	Surface  string `json:"surface"`
	Image    string `json:"image"`
}

type APIResponse struct {
	Response []LeagueResponse `json:"response"`
}

type LeagueResponse struct {
	League  LeagueInfo     `json:"league"`
	Country LeagueCountry  `json:"country"`
	Seasons []LeagueSeason `json:"seasons"`
}

type LeagueCountry struct {
	Name string `json:"name"`
	Code string `json:"code"`
	Flag string `json:"flag"`
}

type LeagueInfo struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
	Logo string `json:"logo"`
}

type LeagueSeason struct {
	Year     int      `json:"year"`
	Start    string   `json:"start"`
	End      string   `json:"end"`
	CURRENT  bool     `json:"current"`
	Coverage Coverage `json:"coverage"`
}

type Coverage struct {
	Fixtures CoverageFixtures `json:"fixtures"`
}

type CoverageFixtures struct {
	Events             bool `json:"events"`
	Lineups            bool `json:"lineups"`
	StatisticsFixtures bool `json:"statistics_fixtures"`
	StatisticsPlayers  bool `json:"statistics_players"`
	Standings          bool `json:"standings"`
	Players            bool `json:"players"`
	TopScorers         bool `json:"top_scorers"`
	TopAssists         bool `json:"top_assists"`
	TopCards           bool `json:"top_cards"`
	Injuries           bool `json:"injuries"`
	Predictions        bool `json:"predictions"`
	Odds               bool `json:"odds"`
}

func (L *LeagueResponse) PrintLeagueresponse() {
	res, err := json.MarshalIndent(L, "", " ")
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return
	}

	fmt.Println(string(res))
}

package bookmaker

import "time"

type MatchOdds struct {
	H2h			[]float32		`json:"h2h"`
}

type SportOddSites struct {
	SiteKey 	string		`json:"site_key"`
	SiteNice	string		`json:"site_nice"`
	LastUpdate	string		`json:"last_update"`
	Odds		MatchOdds	`json:"odds"`
	SitesCount	int			`json:"sites_count"`
}

type SportOdds struct {
	Id				string					`json:"id"`
	SportKey		string					`json:"sport_key"`
	SportNice		string					`json:"sport_nice"`
	Teams			[]string				`json:"teams"`
	CommenceTime	time.Time				`json:"commence_time"`
	HomeTeam		string					`json:"home_team"`
	Sites			[]SportOddSites			`json:"sites"`
}

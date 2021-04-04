package bookmaker

type MatchOdds struct {
	H2h			[]float32	`bson:"h2h" json:"h2h"`
}

type SportOddSites struct {
	SiteKey 	string		`bson:"site_key" json:"site_key"`
	SiteNice	string		`bson:"site_nice" json:"site_nice"`
	LastUpdate	int			`bson:"last_update" json:"last_update"`
	Odds		MatchOdds	`bson:"odds" json:"odds"`
	SitesCount	int			`bson:"sites_count" json:"sites_count"`
}

type SportOddsApiModel struct {
	ApiID			string					`bson:"apiId" json:"id,omitempty"`
	SportKey		string					`bson:"sport_key" json:"sport_key"`
	SportNice		string					`bson:"sport_nice" json:"sport_nice"`
	Teams			[]string				`bson:"teams" json:"teams"`
	CommenceTime	int						`bson:"commence_time" json:"commence_time"`
	HomeTeam		string					`bson:"home_team" json:"home_team"`
	Sites			[]SportOddSites			`bson:"sites" json:"sites"`
}

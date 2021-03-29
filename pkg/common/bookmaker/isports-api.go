package bookmaker

type SportsApi interface {
	FetchSports () ([]Sport, error)
	FetchOdds	(sportType, region, market string) ([]SportOdds, error)
}

package bookmaker

type IBookmakerApi interface {
	FetchSports () ([]Sport, error)
	FetchOdds	(sportType, region, market string) ([]SportOdds, error)
}

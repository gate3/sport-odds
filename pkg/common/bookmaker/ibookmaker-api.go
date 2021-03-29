package bookmaker

type IBookmakerApi interface {
	FetchSports (*[]Sport) ([]Sport, error)
	FetchOdds	(string, string, string, []*SportOdds) error
}

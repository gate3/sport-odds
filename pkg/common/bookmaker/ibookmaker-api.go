package bookmaker

type IBookmakerApi interface {
	FetchSports (*[]SportApiModel) ([]SportApiModel, error)
	FetchOdds	(string, string, string, []*SportOddsApiModel) error
}

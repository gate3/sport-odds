package bookmaker

type IBookmakerApi interface {
	FetchSports (*[]SportApiModel) ([]SportApiModel, error)
	FetchFixtures	(string, string, string, []*SportOddsApiModel) error
}

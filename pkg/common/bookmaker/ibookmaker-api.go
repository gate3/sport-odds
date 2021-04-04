package bookmaker

type IBookmakerApi interface {
	FetchSports 	(*[]SportApiModel) error
	FetchFixtures	(string, string, string, *[]SportOddsApiModel) error
}

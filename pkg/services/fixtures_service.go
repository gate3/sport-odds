package services

import (
	"github.com/gate3/sport-odds/pkg/common/bookmaker"
)

func (s *Services) SaveUpcomingFixtureRecords () (bool, error) {
	odds := new([]bookmaker.SportOddsApiModel)
	err := s.Bookmaker.FetchFixtures("upcoming","uk", "h2h", odds)
	if err != nil {
		return false, err
	}
	_, err = s.Repository.SaveFixtures(odds, s.Db)
	if err != nil {
		return false, err
	}
	return true, nil
}

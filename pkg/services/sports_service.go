package services

import (
	"github.com/gate3/sport-odds/pkg/common/bookmaker"
)

func (s *Services) SaveAllSportsRecords () (bool, error) {
	sp := new([]bookmaker.SportApiModel)
	err := s.Bookmaker.FetchSports(sp)
	if err != nil {
		return false, err
	}

	_, err = s.Repository.SaveSports(sp)
	if err != nil {
		return false, err
	}
	return true, nil
}

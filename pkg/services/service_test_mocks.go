package services

import (
	"github.com/gate3/sport-odds/pkg/common/bookmaker"
	"github.com/gate3/sport-odds/pkg/config"
	"github.com/stretchr/testify/mock"
)

var cfg *config.EnvVariables

func init () {
	cfg = &config.EnvVariables{
		DatabaseUri: "SomeDbUri",
		OddsApiBaseUrl: "SomeApiUrl",
		OddsApiKey: "SomeApiKey",
		DatabaseName: "SomeDbName",
	}
}

type MockRepository struct {
	mock.Mock
}

type MockBookmaker struct {
	mock.Mock
}

func (m *MockRepository) SaveFixtures (sp *[]bookmaker.SportOddsApiModel) ([]interface{}, error) {
	args := m.Called(sp)

	return args.Get(0).([]interface{}), args.Error(1)
}

func (m *MockRepository) SaveSports (sp *[]bookmaker.SportApiModel) ([]interface{}, error) {
	args := m.Called(sp)

	return args.Get(0).([]interface{}), args.Error(1)
}

func (mb *MockBookmaker) FetchFixtures (sportType, region, market string, odds *[]bookmaker.SportOddsApiModel) error {
	args := mb.Called(sportType, region, market, odds)

	return args.Error(0)
}

func (mb *MockBookmaker) FetchSports (sp *[]bookmaker.SportApiModel) error {
	args := mb.Called(sp)

	return args.Error(0)
}

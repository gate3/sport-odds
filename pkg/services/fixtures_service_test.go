package services

import (
	"errors"
	"testing"

	"github.com/gate3/sport-odds/pkg/common/bookmaker"
	"github.com/stretchr/testify/require"
)

func Test_SaveUpcomingFixtureRecords_Runs_Without_Errors (t *testing.T) {

	mockRepo := new(MockRepository)
	odds := new([]bookmaker.SportOddsApiModel)
	mockRepo.On("SaveFixtures", odds).Return([]interface{}{}, nil)

	mockBk := new(MockBookmaker)
	mockBk.On("FetchFixtures", "upcoming","uk", "h2h", odds).Return(nil)

	sr := &Services{
		Repository: mockRepo,
		EnvVars: *cfg,
		Bookmaker: mockBk,
	}

	_, err := sr.SaveUpcomingFixtureRecords()
	if err != nil {
		t.Fail()
	}

	mockRepo.AssertExpectations(t)
	mockBk.AssertExpectations(t)
}

func Test_SaveUpcomingFixtureRecords_Returns_Database_Errors_Correctly (t *testing.T) {
	assertions := require.New(t)
	databaseError := "database Error occurred"

	mockRepo := new(MockRepository)
	odds := new([]bookmaker.SportOddsApiModel)
	mockRepo.On("SaveFixtures", odds).Return([]interface{}{}, errors.New(databaseError))

	mockBk := new(MockBookmaker)
	mockBk.On("FetchFixtures", "upcoming","uk", "h2h", odds).Return(nil)

	sr := &Services{
		Repository: mockRepo,
		EnvVars: *cfg,
		Bookmaker: mockBk,
	}

	_, err := sr.SaveUpcomingFixtureRecords()
	if err != nil && err.Error() != databaseError {
		assertions.Fail("Expected a database error to be thrown")
	}
}

func Test_SaveUpcomingFixtureRecords_Returns_OddsAPI_Errors_Correctly (t *testing.T) {
	assertions := require.New(t)
	mockApiErrorResponse := "Invalid Api Key Provided"

	mockRepo := new(MockRepository)
	odds := new([]bookmaker.SportOddsApiModel)
	mockRepo.On("SaveFixtures", odds).Return([]interface{}{}, nil)

	mockBk := new(MockBookmaker)
	mockBk.On("FetchFixtures", "upcoming","uk", "h2h", odds).Return(errors.New(mockApiErrorResponse))

	sr := &Services{
		Repository: mockRepo,
		EnvVars: *cfg,
		Bookmaker: mockBk,
	}

	_, err := sr.SaveUpcomingFixtureRecords()
	if err != nil && err.Error() != mockApiErrorResponse {
		assertions.Fail("Expected an api error to be thrown")
	}
}

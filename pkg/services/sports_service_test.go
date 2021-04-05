package services

import (
	"errors"
	"testing"

	"github.com/gate3/sport-odds/pkg/common/bookmaker"
	"github.com/stretchr/testify/require"
)

func Test_SaveAllSportsRecords_Runs_Without_Errors (t *testing.T) {
	sp := new([]bookmaker.SportApiModel)

	mockRepo := new(MockRepository)
	mockRepo.On("SaveSports", sp).Return([]interface{}{}, nil)

	mockBk := new(MockBookmaker)
	mockBk.On("FetchSports", sp).Return(nil)

	sr := &Services{
		Repository: mockRepo,
		EnvVars: *cfg,
		Bookmaker: mockBk,
	}

	_, err := sr.SaveAllSportsRecords()
	if err != nil {
		t.Fail()
	}

	mockRepo.AssertExpectations(t)
	mockBk.AssertExpectations(t)
}

func Test_SaveAllSportsRecords_Returns_Database_Errors_Correctly (t *testing.T) {
	assertions := require.New(t)
	sp := new([]bookmaker.SportApiModel)
	mockDatabaseError := "database Error occurred"

	mockRepo := new(MockRepository)
	mockRepo.On("SaveSports", sp).Return([]interface{}{}, errors.New(mockDatabaseError))

	mockBk := new(MockBookmaker)
	mockBk.On("FetchSports", sp).Return(nil)

	sr := &Services{
		Repository: mockRepo,
		EnvVars: *cfg,
		Bookmaker: mockBk,
	}

	_, err := sr.SaveAllSportsRecords()
	if err != nil && err.Error() != mockDatabaseError {
		assertions.Fail("Expected a database error to be thrown")
	}
}

func Test_SaveAllSportsRecords_Returns_API_Errors_Correctly (t *testing.T) {
	assertions := require.New(t)
	sp := new([]bookmaker.SportApiModel)
	mockAPIError := "Invalid api key provided"

	mockRepo := new(MockRepository)
	mockRepo.On("SaveSports", sp).Return([]interface{}{}, nil)

	mockBk := new(MockBookmaker)
	mockBk.On("FetchSports", sp).Return(errors.New(mockAPIError))

	sr := &Services{
		Repository: mockRepo,
		EnvVars: *cfg,
		Bookmaker: mockBk,
	}

	_, err := sr.SaveAllSportsRecords()
	if err != nil && err.Error() != mockAPIError {
		assertions.Fail("Expected a database error to be thrown")
	}
}

package services

import (
	"errors"
	"github.com/gate3/sport-odds/pkg/common/bookmaker"
	"time"
)

type asyncResponse struct {
	status 			bool
	errorObject 	error
}

func (s *Services) SaveUpcomingFixtureRecords (timeoutInSeconds int) (bool, error) {
	c1 := make(chan asyncResponse, 1)
	var response asyncResponse

	go func() {
		time.Sleep(time.Duration(timeoutInSeconds) * time.Second)

		odds := new([]bookmaker.SportOddsApiModel)
		err := s.Bookmaker.FetchFixtures("upcoming","uk", "h2h", odds)
		if err != nil {
			c1 <- asyncResponse{status: false, errorObject: err}
		}
		_, err = s.Repository.SaveFixtures(odds)
		if err != nil {
			c1 <- asyncResponse{status:false, errorObject: err}
		}

		c1 <- asyncResponse{status: true, errorObject: nil}
	}()

	select {
		case res := <- c1:
			response = res
			break
		case <- time.After(time.Duration(timeoutInSeconds + 20) * time.Second): // add an extra 20 seconds to the passed timeout as the limit for the request to timeout
			response = asyncResponse{status: false, errorObject: errors.New("api request timeout exceeded")}
	}

	return response.status, response.errorObject
}

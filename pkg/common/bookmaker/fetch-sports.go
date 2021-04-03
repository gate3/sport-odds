package bookmaker

import (
	"net/http"
)

func (bk BookmakersApi) FetchSports (sports *[]SportApiModel) error {
	fullApiUrl := bk.baseUrl + "sports?apiKey=" + bk.apiToken

	resp, err := http.Get(fullApiUrl)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	err = bk.JsonDecodeApiResponse(resp, sports)
	return err
}


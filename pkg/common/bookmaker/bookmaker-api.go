package bookmaker

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

type BookmakersApi struct {
	baseUrl 	string
	apiToken 	string
}

type ApiResponse struct {
	Success		bool			`json:"success"`
	Data		[]interface{}	`json:"data"`
}

func NewBookmakerApi (baseUrl, apiToken string) (*BookmakersApi, error) {
	if baseUrl == "" {
		return &BookmakersApi{}, errors.New("please provide a base url")
	}
	bookmakerApi := new(BookmakersApi)
	bookmakerApi.baseUrl = baseUrl
	bookmakerApi.apiToken = apiToken
	return bookmakerApi, nil
}

func (bk BookmakersApi) JsonDecodeApiResponse (response *http.Response, result interface{}) error {
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}

	var apiR ApiResponse
	err = json.Unmarshal(body, &apiR)
	if err != nil {
		return err
	}

	if !apiR.Success {
		return errors.New("error fetching resource")
	}

	apiData, _ := json.Marshal(apiR.Data)

	err = json.Unmarshal(apiData, &result)
	if err != nil {
		return err
	}
	return nil
}

package bookmaker

import (
	"net/http"
	"strings"
)

func (bk BookmakersApi) FetchOdds (sportType, region, market string, odds *[]SportOdds) error {
	var stBuilder strings.Builder
	stBuilder.WriteString(bk.baseUrl)
	stBuilder.WriteString("odds?apiKey=" + bk.apiToken)
	stBuilder.WriteString("&sport="+ sportType)
	stBuilder.WriteString("&region="+ region)
	stBuilder.WriteString("&mkt="+market)

	fullApiUrl := stBuilder.String()

	resp, err := http.Get(fullApiUrl)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	err = bk.JsonDecodeApiResponse(resp, odds)
	return err
}

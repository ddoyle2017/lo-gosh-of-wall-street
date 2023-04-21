package warcraft

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	dto "github.com/ddoyle2017/lo-gosh-of-wall-street/internal/dto"
	"github.com/ddoyle2017/lo-gosh-of-wall-street/internal/rest"
)

const OAUTH_ENDPOINT = "https://oauth.battle.net/token"

type BlizzardOAuth struct {
	httpClient   rest.HttpClient
	hostURL      string
	clientID     string
	clientSecret string
}

func (o BlizzardOAuth) NewBlizzardOAuth(httpClient rest.HttpClient, hostURL, clientID, clientSecret string) *BlizzardOAuth {
	return &BlizzardOAuth{
		httpClient:   httpClient,
		hostURL:      hostURL,
		clientID:     clientID,
		clientSecret: clientSecret,
	}
}

func (o BlizzardOAuth) getAuthToken() (token dto.AuthToken, err error) {
	params := url.Values{}
	params.Add("grant_type", "client_credentials")
	query := "?" + params.Encode()

	request, err := http.NewRequest(http.MethodPost, OAUTH_ENDPOINT+query, nil)
	if err != nil {
		return token, err
	}
	request.SetBasicAuth(o.clientID, o.clientSecret)

	response, err := o.httpClient.Do(request)
	if err != nil {
		return token, err
	}
	defer response.Body.Close()

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return token, fmt.Errorf("could parse Blizzard OAuth response body")
	}

	json.Unmarshal(data, &token)
	return token, nil
}

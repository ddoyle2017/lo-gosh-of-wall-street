package warcraft

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	dto "github.com/ddoyle2017/lo-gosh-of-wall-street/internal/dto"
	rest "github.com/ddoyle2017/lo-gosh-of-wall-street/internal/rest"
)

type ClassicAuctionApi struct {
	httpClient rest.HttpClient
	oauth      *BlizzardOAuth
	hostURL    string
}

func NewClassicAuctionApi(httpClient rest.HttpClient, oauth *BlizzardOAuth, hostURL string) *ClassicAuctionApi {
	return &ClassicAuctionApi{
		httpClient: httpClient,
		oauth:      oauth,
		hostURL:    hostURL,
	}
}

func (r ClassicAuctionApi) GetActiveAuctions() (auctions dto.AuctionData, err error) {
	token, err := r.oauth.getAuthToken()
	if err != nil {
		return auctions, fmt.Errorf("could not authenticate to Classic WoW Auctions API")
	}

	request, err := buildRequest(http.MethodGet, r.hostURL+activeAuctionsEndpoint, token)
	if err != nil {
		return auctions, fmt.Errorf("could not build Classic WoW Auctions API request")
	}

	response, err := r.httpClient.Do(request)
	if err != nil {
		return auctions, fmt.Errorf("error when calling Classic WoW Auctions API")
	}
	if response.StatusCode != http.StatusOK {
		return auctions, fmt.Errorf("clasic WoW Auctions API returned a %d", response.StatusCode)
	}
	defer response.Body.Close()

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return auctions, fmt.Errorf("could parse Classic WoW Auctions API response body")
	}

	json.Unmarshal(data, &auctions)
	return auctions, nil
}

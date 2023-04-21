package warcraft

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	dto "github.com/ddoyle2017/lo-gosh-of-wall-street/internal/dto"
	rest "github.com/ddoyle2017/lo-gosh-of-wall-street/internal/rest"
)

type RetailAuctionAPI struct {
	httpClient rest.HttpClient
	hostURL    string
}

func NewRetailAuctionAPI(httpClient rest.HttpClient, hostURL string) *RetailAuctionAPI {
	return &RetailAuctionAPI{
		httpClient: httpClient,
		hostURL:    hostURL,
	}
}

func (r RetailAuctionAPI) GetActiveAuctions() (auctions dto.AuctionData, err error) {
	request, err := buildRequest(http.MethodGet, r.hostURL+activeAuctionsEndpoint)
	if err != nil {
		return auctions, fmt.Errorf("could not build Retail WoW Auctions API request")
	}

	response, err := r.httpClient.Do(request)
	if err != nil {
		return auctions, fmt.Errorf("error when calling Retail WoW Auctions API")
	}
	if response.StatusCode != http.StatusOK {
		return auctions, fmt.Errorf("retail WoW Auctions API returned a %d", response.StatusCode)
	}
	defer response.Body.Close()

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return auctions, fmt.Errorf("could parse Retail WoW Auctions API response body")
	}

	json.Unmarshal(data, &auctions)
	return auctions, nil
}

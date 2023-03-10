package app

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	dto "github.com/ddoyle2017/lo-gosh-of-wall-street/internal/dto"
	rest "github.com/ddoyle2017/lo-gosh-of-wall-street/internal/rest"
)

const activeAuctionsEndpoint string = "/data/wow/connected-realm/1146/auctions"

// const commoditiesEndpoint string = "/data/wow/auctions/commodities"

type WarcraftAuctionAPI interface {
	GetActiveAuctions() (dto.AuctionData, error)
}

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

type ClassicAuctionApi struct {
	httpClient rest.HttpClient
	hostURL    string
}

func NewClassicAuctionApi(httpClient rest.HttpClient, hostURL string) *ClassicAuctionApi {
	return &ClassicAuctionApi{
		httpClient: httpClient,
		hostURL:    hostURL,
	}
}

func (r ClassicAuctionApi) GetActiveAuctions() (auctions dto.AuctionData, err error) {
	request, err := buildRequest(http.MethodGet, r.hostURL+activeAuctionsEndpoint)
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

// Builds out a Blizzard API request + headers, including the correct authorization
func buildRequest(method string, url string) (request *http.Request, err error) {
	request, err = http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}

	// TODO get OAuth token from Blizzard API auth endpoint

	request.Header.Add("Accept", `application/json`)
	request.Header.Add("Authorization", "")
	return request, nil
}

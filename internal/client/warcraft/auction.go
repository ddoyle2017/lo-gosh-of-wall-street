package warcraft

import (
	"net/http"

	dto "github.com/ddoyle2017/lo-gosh-of-wall-street/internal/dto"
)

const activeAuctionsEndpoint string = "/data/wow/connected-realm/1146/auctions"

// const commoditiesEndpoint string = "/data/wow/auctions/commodities"

type AuctionAPI interface {
	GetActiveAuctions() (dto.AuctionData, error)
}

// Builds out a Blizzard API request + headers, including the correct authorization
func buildRequest(method string, url string, token dto.AuthToken) (request *http.Request, err error) {
	request, err = http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}

	request.Header.Add("Accept", `application/json`)
	request.Header.Add("Authorization", "Bearer "+token.AccessToken)
	return request, nil
}

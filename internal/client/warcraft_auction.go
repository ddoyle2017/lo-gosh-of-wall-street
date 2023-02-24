package app

import (
	"encoding/json"
	"io/ioutil"
	"log"

	rest "github.com/ddoyle2017/lo-gosh-of-wall-street/internal/rest"
)

type AuctionLength string

const (
	SHORT     AuctionLength = "SHORT"
	MEDIUM    AuctionLength = "MEDIUM"
	LONG      AuctionLength = "LONG"
	VERY_LONG AuctionLength = "VERY_LONG"
)

type Modifier struct {
	Type_ uint
	Value uint
}

type Item struct {
	Id         uint
	BonusLists []uint
	Context    uint // The context an item was obtained in, e.g. mythic raid, dungeon, quest, etc.
	Modifiers  []Modifier
}

type Auction struct {
	Id        uint64
	Buyout    uint64
	Bid       uint64
	Item      Item
	Quantity  uint
	UnitPrice uint64
	TimeLeft  AuctionLength
}

type AuctionData struct {
	Auctions []Auction
}

type WarcraftAuctionApi interface {
	GetActiveAuctions() AuctionData
}

type RetailAuctionApi struct {
	httpClient       rest.HttpClient
	auctionsEndpoint string
}

func NewRetailAuctionApi(httpClient rest.HttpClient, auctionsEndpoint string) *RetailAuctionApi {
	return &RetailAuctionApi{
		httpClient:       httpClient,
		auctionsEndpoint: auctionsEndpoint,
	}
}

func (r RetailAuctionApi) GetActiveAuctions() (auctions AuctionData) {
	response, err := r.httpClient.Get(r.auctionsEndpoint)
	handleError(err)

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	json.Unmarshal(data, &auctions)
	return
}

type ClassicAuctionApi struct {
	httpClient       rest.HttpClient
	auctionsEndpoint string
}

func NewClassicAuctionApi(httpClient rest.HttpClient, auctionsEndpoint string) *ClassicAuctionApi {
	return &ClassicAuctionApi{
		httpClient:       httpClient,
		auctionsEndpoint: auctionsEndpoint,
	}
}

func (r ClassicAuctionApi) GetActiveAuctions() (auctions AuctionData) {
	response, err := r.httpClient.Get(r.auctionsEndpoint)
	handleError(err)

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	json.Unmarshal(data, &auctions)
	return
}

func handleError(err error) {
	if err != nil {
		log.Fatal("World of Warcraft API returned an error")
	}
}

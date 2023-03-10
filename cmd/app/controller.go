package main

import (
	"context"
	"fmt"

	client "github.com/ddoyle2017/lo-gosh-of-wall-street/internal/client"
	dto "github.com/ddoyle2017/lo-gosh-of-wall-street/internal/dto"
)

// The Controller is responsible for business logic. Taking data and transforming it into something we want to return
// to our API callers.
type Controller interface {
	ListTopAuctions(context.Context, int) ([]dto.Auction, error)
}

type controller struct {
	auctionAPI client.WarcraftAuctionAPI
}

func NewController(a client.WarcraftAuctionAPI) Controller { return controller{auctionAPI: a} }

func (c controller) ListTopAuctions(ctx context.Context, ranking int) (auctions []dto.Auction, err error) {

	// Get auction data
	auctionData, err := c.auctionAPI.GetActiveAuctions()
	if err != nil {
		return nil, fmt.Errorf("WoW API call failed")
	}
	auctions = auctionData.Auctions

	return auctions, nil
}

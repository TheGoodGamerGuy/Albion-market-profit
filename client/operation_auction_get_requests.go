package client

import (
	"encoding/json"

	"github.com/ao-data/albiondata-client/lib"
	"github.com/ao-data/albiondata-client/log"
)

func contains(s []*lib.MarketOrder, what int) bool {
	for _, v := range s {
		if v.ID == what {
			return true
		}
	}

	return false
}

type operationAuctionGetRequestsResponse struct {
	MarketOrders []string `mapstructure:"0"`
}

func (op operationAuctionGetRequestsResponse) Process(state *albionState) {
	log.Debug("Got response to AuctionGetOffers operation...")

	if !state.IsValidLocation() {
		return
	}

	var orders []*lib.MarketOrder

	for _, v := range op.MarketOrders {
		order := &lib.MarketOrder{}
		err := json.Unmarshal([]byte(v), order)

		if err != nil {
			log.Errorf("Problem converting market order to internal struct: %v", err)
		}

		order.LocationID = state.LocationId
		if !contains(allRequests, order.ID) {
			allRequests = append(allRequests, order)
		}
		orders = append(orders, order)
	}

	if len(orders) < 1 {
		return
	}

	log.Infof("Found %d new market requests", len(orders))
	log.Infof("Total length: %d", len(allRequests))

	// getBestCraftingProfit()
	getBestProfit()

}

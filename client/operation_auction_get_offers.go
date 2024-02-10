package client

import (
	"encoding/json"

	"github.com/ao-data/albiondata-client/lib"
	"github.com/ao-data/albiondata-client/log"
	"time"
	"google.golang.org/api/sheets/v4"
	"fmt"
)

type operationAuctionGetOffers struct {
	Category         string   `mapstructure:"1"`
	SubCategory      string   `mapstructure:"2"`
	Quality          string   `mapstructure:"3"`
	Enchantment      uint32   `mapstructure:"4"`
	EnchantmentLevel string   `mapstructure:"8"`
	ItemIds          []uint16 `mapstructure:"6"`
	MaxResults       uint32   `mapstructure:"9"`
	IsAscendingOrder bool     `mapstructure:"12"`
}

func (op operationAuctionGetOffers) Process(state *albionState) {
	log.Debug("Got AuctionGetOffers operation...")
}

type operationAuctionGetOffersResponse struct {
	MarketOrders []string `mapstructure:"0"`
}

func extractbuyFields(orders []*lib.MarketOrder) [][]interface{} {
	var result [][]interface{}

	for _, order := range orders {
		priceDivided := order.Price / 10000
		currentTime := time.Now().UTC().Unix()

		orderSlice := []interface{}{
			order.ID,
			order.ItemID,
			order.GroupTypeId,
			order.LocationID,
			order.QualityLevel,
			order.EnchantmentLevel,
			priceDivided,
			order.Amount,
			order.AuctionType,
			currentTime,
		}

		result = append(result, orderSlice)
	}

	return result
}

func (op operationAuctionGetOffersResponse) Process(state *albionState) {
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
		orders = append(orders, order)
	}

	if len(orders) < 1 {
		return
	}


	log.Infof("Sending %d market offers to ingest", len(orders))

	record := extractbuyFields(orders)

	spreadsheetId := "1q34v4KwaJAFK7-L1pubIIuKm-qBPIHu1cFC_k9rqWWg"
	name := "instabuy"
	valueInputOption := "USER_ENTERED"
	insertDataOption := "INSERT_ROWS"

	rb := &sheets.ValueRange{
		Values: record,
	}
	response, err := globalSheetService.Spreadsheets.Values.Append(spreadsheetId, name, rb).ValueInputOption(valueInputOption).InsertDataOption(insertDataOption).Context(ctx).Do()
	if err != nil || response.HTTPStatusCode != 200 {
			fmt.Println(err)
	}
}

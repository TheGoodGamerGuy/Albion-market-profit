package client

import (
	"encoding/json"

	"github.com/ao-data/albiondata-client/lib"
	"github.com/ao-data/albiondata-client/log"
	"google.golang.org/api/sheets/v4"
	"fmt"
)

type operationAuctionGetRequestsResponse struct {
	MarketOrders []string `mapstructure:"0"`
}

type MarketOrder struct {
    ID               int    `json:"Id"`
    ItemID           string `json:"ItemTypeId"`
    GroupTypeId      string `json:"ItemGroupTypeId"`
    LocationID       int    `json:"LocationId"`
    QualityLevel     int    `json:"QualityLevel"`
    EnchantmentLevel int    `json:"EnchantmentLevel"`
    Price            int    `json:"UnitPriceSilver"`
    Amount           int    `json:"Amount"`
    AuctionType      string `json:"AuctionType"`
    // Expires          string `json:"Expires"`
}

func extractFields(orders []*lib.MarketOrder) [][]interface{} {
	var result [][]interface{}

	// Iterate through each order
	for _, order := range orders {
		// Divide the Price field by 10000
		priceDivided := order.Price / 10000

		// Create a slice with the desired fields, including the divided price
		orderSlice := []interface{}{
			order.ID,
			order.ItemID,
			order.GroupTypeId,
			order.LocationID,
			order.QualityLevel,
			order.EnchantmentLevel,
			priceDivided, // Use the divided price here
			order.Amount,
			order.AuctionType,
		}

		// Append the slice to the result
		result = append(result, orderSlice)
	}

	return result
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
		orders = append(orders, order)
	}

	if len(orders) < 1 {
		return
	}

	// upload := lib.MarketUpload{
	// 	Orders: orders,
	// }

	log.Infof("Sending %d market requests to ingest", len(orders))
	// sendMsgToPublicUploaders(upload, lib.NatsMarketOrdersIngest, state)
	record := extractFields(orders)

	spreadsheetId := "1q34v4KwaJAFK7-L1pubIIuKm-qBPIHu1cFC_k9rqWWg"
	name := "test"
	valueInputOption := "USER_ENTERED"
	insertDataOption := "INSERT_ROWS"

	rb := &sheets.ValueRange{
		Values: record,
	}
	response, err := globalSheetService.Spreadsheets.Values.Append(spreadsheetId, name, rb).ValueInputOption(valueInputOption).InsertDataOption(insertDataOption).Context(ctx).Do()
	if err != nil || response.HTTPStatusCode != 200 {
			fmt.Println(err)
	}
	// for i := 0; i < len(orders); i++ {
	// 	log.Info(orders[i])
	// }
	// log.Print(orders[0].Price/10000)
}

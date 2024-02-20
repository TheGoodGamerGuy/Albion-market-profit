package client

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

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

func extractFields(orders []*lib.MarketOrder) [][]interface{} {
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

func getBestProfit() {
	for _, buyOrder := range allRequests {
		var name = buyOrder.ItemID
		var quality = buyOrder.QualityLevel
		for _, sellOrder := range allOffers {
			if sellOrder.ItemID == name && sellOrder.QualityLevel >= quality {
				var profit = ((buyOrder.Price - (buyOrder.Price * 4 / 100)) - sellOrder.Price) / 10000
				if profit > 0 {
					for _, item := range itemseverything {
						if item.UniqueName == name {
							name = item.LocalizedNames.EN_US
							break
						}
					}
					profit := strconv.Itoa(profit)
					var quality string
					switch qualityNum := sellOrder.QualityLevel; qualityNum {
					case 1:
						quality = "normal"
					case 2:
						quality = "good"
					case 3:
						quality = "outstanding"
					case 4:
						quality = "excellent"
					case 5:
						quality = "masterpiece"
					}
					var enchantmentLevel string = strconv.Itoa(sellOrder.EnchantmentLevel)
					// log.Info("buy price: " + strconv.Itoa(sellOrder.Price/10000))
					// log.Info("sell price: " + strconv.Itoa(buyOrder.Price/10000))
					// log.Info(profit + " silver profit, buy " + name + " with " + quality + " quality and enchantment level " + enchantmentLevel)
					var buyLocation string
					var sellLocation string
					for _, location := range world {
						if location.Index == strconv.Itoa(sellOrder.LocationID) {
							buyLocation = location.UniqueName
						}
						if location.Index == strconv.Itoa(buyOrder.LocationID) {
							sellLocation = location.UniqueName
						}
						if sellLocation == "Caerleon" {
							sellLocation = "Black Market"
						}
					}
					fmt.Println("")
					fmt.Println("buy at " + buyLocation + ", price: " + strconv.Itoa(sellOrder.Price/10000))
					fmt.Println("sell at " + sellLocation + ", price: " + strconv.Itoa(buyOrder.Price/10000))
					fmt.Println("Profit: " + profit + " silver , buy " + name + " with " + quality + " quality and enchantment level " + enchantmentLevel)
					fmt.Println("")
				}
			}
		}
	}
}

// func GetCraftingProfit() {
// 	for _, buyOrder := range allRequests {
// 		var name = buyOrder.ItemID
// 		var quality = buyOrder.QualityLevel

// 	}
// }

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
		// log.Info(order.Price)
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

	getBestProfit()

	// record := extractFields(orders)
	// spreadsheetId := "1q34v4KwaJAFK7-L1pubIIuKm-qBPIHu1cFC_k9rqWWg"
	// name := "instasell"
	// valueInputOption := "USER_ENTERED"
	// insertDataOption := "INSERT_ROWS"

	// rb := &sheets.ValueRange{
	// 	Values: record,
	// }
	// response, err := globalSheetService.Spreadsheets.Values.Append(spreadsheetId, name, rb).ValueInputOption(valueInputOption).InsertDataOption(insertDataOption).Context(ctx).Do()
	// if err != nil || response.HTTPStatusCode != 200 {
	// 	log.Error(err)
	// }
}

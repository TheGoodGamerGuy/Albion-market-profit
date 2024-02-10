package main

// import (
//         "context"
//         "encoding/json"
//         "fmt"
//         "log"
//         "net/http"

//         "golang.org/x/oauth2"
//         "golang.org/x/oauth2/google"
//         "google.golang.org/api/option"
//         "google.golang.org/api/sheets/v4"
// )


// type MarketOrder struct {
//     ID               int    `json:"Id"`
//     ItemID           string `json:"ItemTypeId"`
//     GroupTypeId      string `json:"ItemGroupTypeId"`
//     LocationID       int    `json:"LocationId"`
//     QualityLevel     int    `json:"QualityLevel"`
//     EnchantmentLevel int    `json:"EnchantmentLevel"`
//     Price            int    `json:"UnitPriceSilver"`
//     Amount           int    `json:"Amount"`
//     AuctionType      string `json:"AuctionType"`
//     // Expires          string `json:"Expires"`
// }

// func extractFields(orders []MarketOrder) [][]interface{} {
// 	var result [][]interface{}

// 	// Iterate through each order
// 	for _, order := range orders {
// 		// Create a slice with the desired fields
// 		orderSlice := []interface{}{
// 			order.ID,
// 			order.ItemID,
// 			order.GroupTypeId,
// 			order.LocationID,
// 			order.QualityLevel,
// 			order.EnchantmentLevel,
// 			order.Price,
// 			order.Amount,
// 			order.AuctionType,
// 		}

// 		// Append the slice to the result
// 		result = append(result, orderSlice)
// 	}

// 	return result
// }

// func getClient(config *oauth2.Config) *http.Client {
// 	jsonData := []byte(`{"access_token":"ya29.a0AfB_byA-yynquN8gNrFpr2e2idUD8QRFON1xCUzTdeXLAPxItxpiRLxW4bTWNU0XVmsv5YKj4HtFckjy-vv6lvSc8Ltu3eM6prnmPrwI2iVukTIF_Qzew9r4U5GI91RQZX83Kh9NmZM__zFrUNUWM-RLoyOh6wA0LGpkaCgYKAWoSARISFQHGX2MiUffLQSnF88IVFOd6D72p2A0171","token_type":"Bearer","refresh_token":"1//0cHBoqnYRsihyCgYIARAAGAwSNwF-L9IrtJZMKvpHrhJ3iElHXfcQUIC2gm3gUQg2ROt4KGyGYKFxlI0jMy2Qm9FdZ6tI3ApE-hA","expiry":"2024-01-23T21:32:52.9548704+02:00"}`)
//         tok := &oauth2.Token{}
//         err := json.Unmarshal(jsonData, tok)
//         if err != nil {
//                 fmt.Println("error")
//         }
//         return config.Client(context.Background(), tok)
// }






// func main() {
//         ctx := context.Background()

//         b := []byte(`{"installed":{"client_id":"907770711446-4hd5je7agnh8gk0sno1acl0v67i33nic.apps.googleusercontent.com","project_id":"albionshit","auth_uri":"https://accounts.google.com/o/oauth2/auth","token_uri":"https://oauth2.googleapis.com/token","auth_provider_x509_cert_url":"https://www.googleapis.com/oauth2/v1/certs","client_secret":"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","redirect_uris":["http://localhost"]}}`)

//         config, err := google.ConfigFromJSON(b, "https://www.googleapis.com/auth/spreadsheets")
//         if err != nil {
//                 log.Fatalf("Unable to parse client secret file to config: %v", err)
//         }
//         client := getClient(config)

//         srv, err := sheets.NewService(ctx, option.WithHTTPClient(client))
//         if err != nil {
//                 log.Fatalf("Unable to retrieve Sheets client: %v", err)
//         }




//         spreadsheetId := "1q34v4KwaJAFK7-L1pubIIuKm-qBPIHu1cFC_k9rqWWg"
//         name := "test"
//         valueInputOption := "USER_ENTERED"
//         insertDataOption := "INSERT_ROWS"

// 	orders := []MarketOrder{
// 		{1, "item1", "group1", 123, 5, 2, 100, 10, "Auction"},
// 		{2, "item2", "group2", 456, 3, 1, 200, 5, "Auction"},
// 		{3, "item3", "group3", 789, 7, 0, 150, 8, "Auction"},
// 	}
// 	record := extractFields(orders)

//         rb := &sheets.ValueRange{
//                 Values: record,
//         }
//         response, err := srv.Spreadsheets.Values.Append(spreadsheetId, name, rb).ValueInputOption(valueInputOption).InsertDataOption(insertDataOption).Context(ctx).Do()
//         if err != nil || response.HTTPStatusCode != 200 {
//                 fmt.Println(err)
//         }
// }
package client

import (
	"fmt"
	"strconv"

	"github.com/ao-data/albiondata-client/lib"
	"github.com/ao-data/albiondata-client/log"
)

var allRecipes []AllItemRecipes

type AllItemRecipes struct {
	name        string
	itemRecipes ItemRecipes
}

type ItemRecipes struct {
	allRecipes []Recipe
}

// type AllRecipes struct {
// 	recipe []Recipe
// }

type Recipe struct {
	resource []Resource
}

type Resource struct {
	name        string
	amount      int
	singlePrice int
	fullPrice   int
}

type SellOrders struct {
	sellOrder lib.MarketOrder
}

func getBestProfit() {
	for _, buyOrder := range allRequests {
		var preciseName = buyOrder.ItemID
		var quality = buyOrder.QualityLevel
		for _, sellOrder := range allOffers {
			if sellOrder.ItemID == preciseName && sellOrder.QualityLevel >= quality {
				var profit = ((buyOrder.Price - (buyOrder.Price * 4 / 100)) - sellOrder.Price) / 10000
				if profit-40000 > 0 {
					profit := strconv.Itoa(profit)

					// item name
					var name = buyOrder.GroupTypeId
					for _, item := range items {
						if item[0] == name {
							name = item[1]
							break
						}
					}

					// quality
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

					// enchantments
					var enchantmentLevel string = strconv.Itoa(sellOrder.EnchantmentLevel)

					// location
					var buyLocation string
					var sellLocation string
					for _, location := range worlds {
						if location[0] == strconv.Itoa(sellOrder.LocationID) {
							buyLocation = location[1]
						}
						if location[0] == strconv.Itoa(buyOrder.LocationID) {
							sellLocation = location[1]
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
		// for _, sellOrderRecipe := range allRecipes {
		// 	if sellOrderRecipe.name == buyOrder.ItemID {
		// 		for _, recipe := range sellOrderRecipe.itemRecipes.allRecipes {
		// 			resourcelength := 0
		// 			for _, resource := range recipe.resource {
		// 				for _, sellOrder := range allOffers {
		// 					if sellOrder.ItemID == resource.name {
		// 						for _, preciseSellOrder := range allOffers {

		// 							if preciseSellOrder.ItemID == resource.name {

		// 							}
		// 						}
		// 						// resourcelength++
		// 						// break
		// 					}
		// 				}
		// 			}
		// 			if resourcelength == len(recipe.resource) {
		// 				fmt.Println("")
		// 				fmt.Println(sellOrderRecipe.name)
		// 				fmt.Println(recipe.resource)
		// 				fmt.Println("")
		// 			}
		// 		}
		// 	}
		// }
	}
}

func getBestCraftingProfit() {
	allitems := itemseverything.Items
	for _, buyOrder := range allRequests {
		var name = buyOrder.GroupTypeId
		for _, weapon := range allitems.Weapon {
			if weapon.Uniquename == name {
				var allItemRecipe AllItemRecipes
				allItemRecipe.name = name
				var itemRecipe ItemRecipes
				switch first := weapon.Craftingrequirements.(type) {
				case []interface{}:
					for _, secondfa := range first {
						var recipe Recipe
						thirdfa := secondfa.(map[string]interface{})
						switch second := thirdfa["craftresource"].(type) {
						case []interface{}:
							for _, third := range second {
								var resource Resource
								fourth := third.(map[string]interface{})
								resource.name = fourth["@uniquename"].(string)
								amount, err := strconv.Atoi(fourth["@count"].(string))
								if err != nil {
									log.Error(err)
								}
								resource.amount = amount
								recipe.resource = append(recipe.resource, resource)
							}
						case map[string]interface{}:
							var resource Resource
							resource.name = second["@uniquename"].(string)
							amount, err := strconv.Atoi(second["@count"].(string))
							if err != nil {
								log.Error(err)
							}
							resource.amount = amount
							recipe.resource = append(recipe.resource, resource)
						}
						itemRecipe.allRecipes = append(itemRecipe.allRecipes, recipe)
					}
				case map[string]interface{}:
					var recipe Recipe
					switch second := first["craftresource"].(type) {
					case []interface{}:
						for _, third := range second {
							var resource Resource
							fourth := third.(map[string]interface{})
							resource.name = fourth["@uniquename"].(string)
							amount, err := strconv.Atoi(fourth["@count"].(string))
							if err != nil {
								log.Error(err)
							}
							resource.amount = amount
							recipe.resource = append(recipe.resource, resource)
						}
					case map[string]interface{}:
						var resource Resource
						resource.name = second["@uniquename"].(string)
						amount, err := strconv.Atoi(second["@count"].(string))
						if err != nil {
							log.Error(err)
						}
						resource.amount = amount
						recipe.resource = append(recipe.resource, resource)
					}
					itemRecipe.allRecipes = append(itemRecipe.allRecipes, recipe)
				}
				allItemRecipe.itemRecipes = itemRecipe
				allRecipes = append(allRecipes, allItemRecipe)
				break
			}
		}
		// 	for _, transweapon := range allitems.Transformationweapon {
		// 		if transweapon.Uniquename == name {
		// 			var allItemRecipe AllItemRecipes
		// 			allItemRecipe.name = name
		// 			var itemRecipe ItemRecipes
		// 			switch first := transweapon.Craftingrequirements.(type) {
		// 			case []interface{}:
		// 				for _, secondfa := range first {
		// 					var recipe Recipe
		// 					thirdfa := secondfa.(map[string]interface{})
		// 					switch second := thirdfa["craftresource"].(type) {
		// 					case []interface{}:
		// 						for _, third := range second {
		// 							var resource Resource
		// 							fourth := third.(map[string]interface{})
		// 							resource.name = fourth["@uniquename"].(string)
		// 							resource.amount = fourth["@count"].(string)
		// 							recipe.resource = append(recipe.resource, resource)
		// 						}
		// 					case map[string]interface{}:
		// 						var resource Resource
		// 						resource.name = second["@uniquename"].(string)
		// 						resource.amount = second["@count"].(string)
		// 						recipe.resource = append(recipe.resource, resource)
		// 					}
		// 					itemRecipe.allRecipes = append(itemRecipe.allRecipes, recipe)
		// 				}
		// 			case map[string]interface{}:
		// 				var recipe Recipe
		// 				switch second := first["craftresource"].(type) {
		// 				case []interface{}:
		// 					for _, third := range second {
		// 						var resource Resource
		// 						fourth := third.(map[string]interface{})
		// 						resource.name = fourth["@uniquename"].(string)
		// 						resource.amount = fourth["@count"].(string)
		// 						recipe.resource = append(recipe.resource, resource)
		// 					}
		// 				case map[string]interface{}:
		// 					var resource Resource
		// 					resource.name = second["@uniquename"].(string)
		// 					resource.amount = second["@count"].(string)
		// 					recipe.resource = append(recipe.resource, resource)
		// 				}
		// 				itemRecipe.allRecipes = append(itemRecipe.allRecipes, recipe)
		// 			}
		// 			allItemRecipe.itemRecipes = itemRecipe
		// 			allRecipes = append(allRecipes, allItemRecipe)
		// 			break
		// 		}
		// 	}
	}
	// fmt.Println(allRecipes)
}

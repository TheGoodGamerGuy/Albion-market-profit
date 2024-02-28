package client

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/ao-data/albiondata-client/lib"
	"github.com/ao-data/albiondata-client/log"
	// "golang.org/x/oauth2"
	// "golang.org/x/oauth2/google"
	// "google.golang.org/api/option"
	// "google.golang.org/api/sheets/v4"
)

// var version string
var allRequests []*lib.MarketOrder
var allOffers []*lib.MarketOrder

var itemseverything Itemseverything

// Client struct base
type Client struct {
}

// NewClient return a new Client instance
func NewClient(_version string) *Client {
	// version = _version
	return &Client{}
}

// Run starts client settings and run
func (client *Client) Run() error {

	log.Info("Starting...")

	log.Info("Downloading items.json")
	url := "https://raw.githubusercontent.com/ao-data/ao-bin-dumps/master/items.json"
	Response, err := http.Get(url)
	if err != nil {
		fmt.Println("Error making GET request:", err)
	}
	defer Response.Body.Close()
	log.Info("Reading items.json")
	body, err := io.ReadAll(Response.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
	}
	json.Unmarshal([]byte(body), &itemseverything)

	log.Info("items.json ready")

	fmt.Println(itemseverything.XML)

	apw := newAlbionProcessWatcher()
	return apw.run()
}

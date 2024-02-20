package client

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"time"

	"github.com/ao-data/albiondata-client/lib"
	"github.com/ao-data/albiondata-client/log"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

// var version string
var globalSheetService *sheets.Service
var ctx context.Context
var allRequests []*lib.MarketOrder
var allOffers []*lib.MarketOrder

// var itemseverything map[string]interface{}
var itemseverything []Items

type Items struct {
	UniqueName     string
	LocalizedNames LocalizedNames
}

type LocalizedNames struct {
	EN_US string `json:"EN-US"`
}

// Client struct base
type Client struct {
}

// NewClient return a new Client instance
func NewClient(_version string) *Client {
	// version = _version
	return &Client{}
}

func getClient(config *oauth2.Config) *http.Client {
	jsonData := []byte(`{"access_token":"ya29.a0AfB_byDR7V5hUdZp1o5YXpoTcv2_hQdQ9TukqwXlDsd6YKFP5cpxZPsEIlr0yspk3FHJgIKq1hwbrMdsLyELp06omqLSJcw8SiMOWaMfmqmH3F-4kdNgv8j0z7kz2sqrQEaj4sTAAZGdmtZaF59HBRkbIOgNK9IlHLdaaCgYKAfwSARISFQHGX2Mi6wectlPuOCfGG65AE4LiBA0171","token_type":"Bearer","refresh_token":"1//0cyOyL13V_OqWCgYIARAAGAwSNwF-L9IrXKAlXM5IaYn8BJTUkZEh5tLFVEXftys5uzcVhfk1yhC_W_M5zex-_rVn0xU1W1KTOWI","expiry":"2024-02-07T15:11:11.772736+02:00"}`)
	tok := &oauth2.Token{}
	err := json.Unmarshal(jsonData, tok)
	if err != nil {
		fmt.Println("error")
	}
	return config.Client(context.Background(), tok)
}

func googleSheets() (*sheets.Service, error) {
	ctx := context.Background()

	b := []byte(`{"installed":{"client_id":"907770711446-4hd5je7agnh8gk0sno1acl0v67i33nic.apps.googleusercontent.com","project_id":"albionshit","auth_uri":"https://accounts.google.com/o/oauth2/auth","token_uri":"https://oauth2.googleapis.com/token","auth_provider_x509_cert_url":"https://www.googleapis.com/oauth2/v1/certs","client_secret":"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","redirect_uris":["http://localhost"]}}`)

	config, err := google.ConfigFromJSON(b, "https://www.googleapis.com/auth/spreadsheets")
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
	}
	client := getClient(config)

	srv, err := sheets.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		log.Fatalf("Unable to retrieve Sheets client: %v", err)
		return nil, err
	}

	globalSheetService = srv

	return srv, nil
}

// Run starts client settings and run
func (client *Client) Run() error {
	// log.Infof("Starting Albion Data Client, version: %s", version)
	// log.Info("This is a third-party application and is in no way affiliated with Sandbox Interactive or Albion Online.")
	// log.Info("Additional parameters can listed by calling this file with the -h parameter.")
	log.Info("Starting...")

	log.Info("Setting up Google Sheets")
	sheetsService, err := googleSheets()
	if err != nil {
		log.Fatalf("Error setting up Google Sheets: %v", err)
		return err
	}
	log.Info(sheetsService)

	//JSON
	log.Info("Reading items.json")

	//Open JSON file
	// f, err := os.Open("items.json")
	// if err != nil {
	// 	log.Error(err)
	// }
	// defer f.Close()
	// byteValue, _ := io.ReadAll(f)
	// json.Unmarshal([]byte(byteValue), &itemseverything)

	//Get JSON from URL
	url := "https://raw.githubusercontent.com/ao-data/ao-bin-dumps/master/formatted/items.json"

	spaceClient := http.Client{
		Timeout: time.Second * 2, // Timeout after 2 seconds
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("User-Agent", "spacecount-tutorial")

	res, getErr := spaceClient.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	byteValue, _ := io.ReadAll(res.Body)
	json.Unmarshal([]byte(byteValue), &itemseverything)

	log.Info("items.json ready")

	// log.Info(itemseverything[0].LocalizedNames.EN_US)

	// f, err := os.Open("items.json")
	// if err != nil {
	// 	log.Error(err)
	// }
	// defer f.Close()
	// bytesValue, _ := io.ReadAll(f)
	// var objmap []map[string]interface{}
	// if err := json.Unmarshal(bytesValue, &objmap); err != nil {
	// 	log.Fatal(err)
	// }
	// log.Info(objmap[0]["UniqueName"])
	// ConfigGlobal.setupDebugEvents()
	// ConfigGlobal.setupDebugOperations()

	// createDispatcher()

	if ConfigGlobal.Offline {
		processOffline(ConfigGlobal.OfflinePath)
	} else {
		apw := newAlbionProcessWatcher()
		return apw.run()
	}

	return nil
}

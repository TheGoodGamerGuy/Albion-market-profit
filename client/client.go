package client

import (
	"github.com/ao-data/albiondata-client/log"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

var version string
var globalSheetService *sheets.Service
var ctx context.Context

//Client struct base
type Client struct {
}

//NewClient return a new Client instance
func NewClient(_version string) *Client {
	version = _version
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



//Run starts client settings and run
func (client *Client) Run() error {
	// log.Infof("Starting Albion Data Client, version: %s", version)
	// log.Info("This is a third-party application and is in no way affiliated with Sandbox Interactive or Albion Online.")
	// log.Info("Additional parameters can listed by calling this file with the -h parameter.")
	log.Info("Starting...")

	sheetsService, err := googleSheets()
	if err != nil {
		log.Fatalf("Error setting up Google Sheets: %v", err)
		return err
	}

	fmt.Println(sheetsService)

	ConfigGlobal.setupDebugEvents()
	ConfigGlobal.setupDebugOperations()

	// createDispatcher()

	if ConfigGlobal.Offline {
		processOffline(ConfigGlobal.OfflinePath)
	} else {
		apw := newAlbionProcessWatcher()
		return apw.run()
	}


	return nil
}
package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/ao-data/albiondata-client/log"
)

var world []World

type World struct {
	Index      string
	UniqueName string
}

func main() {
	log.Info("reading world.json")
	worldURL := "https://raw.githubusercontent.com/ao-data/ao-bin-dumps/master/formatted/world.json"

	// Make GET request
	worldResponse, err := http.Get(worldURL)
	if err != nil {
		fmt.Println("Error making GET request:", err)
	}
	defer worldResponse.Body.Close()

	// Read the response body
	worldbody, err := io.ReadAll(worldResponse.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
	}

	json.Unmarshal([]byte(worldbody), &world)

	// log.Info(world[0].Index)
	log.Info(world)
}

package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// Settings -- object containing the configuration information for the bot
var Settings struct {
	ConsumerKey       string
	ConsumerSecret    string
	AccessKey         string
	AccessToken       string
	TweetInterests    []string
	FollowerThreshold int
}

// ReadSettings -- information is read from config.json and stores it in the settings struct
func ReadSettings() {

	configFile, err := os.Open("./config.json")
	if err != nil {
		fmt.Println("Problem opening config.json", err)
	}
	jsonParser := json.NewDecoder(configFile)
	if err = jsonParser.Decode(&Settings); err != nil {
		fmt.Println("parsing config file", err)
	}

	if err != nil {
		fmt.Println("Problem with unmarshal", err)
	}

	return // Naked return, information is being placed directly into the settings struct by jsonParser.Decode()
}

package main

import (
	"log"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

func GsheetConfig() *oauth2.Config {
	b, err := Asset("client_secret.json")
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
	}
	config, err := google.ConfigFromJSON(b, "https://www.googleapis.com/auth/spreadsheets")
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
	}
	return config
}

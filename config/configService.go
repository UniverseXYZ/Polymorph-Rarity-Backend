package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"rarity-backend/structs"
)

func NewConfigService(configPath string) *structs.ConfigService {
	jsonFile, err := os.Open(configPath)

	if err != nil {
		log.Fatal("Missing polymorphs config file")
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var service structs.ConfigService

	json.Unmarshal(byteValue, &service)

	return &service
}

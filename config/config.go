package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

type AttributeSet struct {
	Name string   `json:"name"`
	Sets []string `json:"sets"`
}

type ComboSet struct {
	Name        string `json:"name"`
	TraitsCount int    `json:"traitsCount"`
}

type ConfigService struct {
	Type        []string       `json:"type"`
	Character   []string       `json:"character"`
	Background  []string       `json:"background"`
	Footwear    []AttributeSet `json:"footwear"`
	Pants       []AttributeSet `json:"pants"`
	Torso       []AttributeSet `json:"torso"`
	Eyewear     []AttributeSet `json:"eyewear"`
	Headwear    []AttributeSet `json:"headwear"`
	WeaponRight []AttributeSet `json:"weaponright"`
	WeaponLeft  []AttributeSet `json:"weaponleft"`
}

type SetWithColors struct {
	Name           string
	Colors         []string
	TraitsNumber   float64
	NonColorTraits float64
}

func NewConfigService(configPath string) *ConfigService {
	jsonFile, err := os.Open(configPath)
	// if we os.Open returns an error then handle it
	if err != nil {
		log.Fatal("Missing polymorphs config file")
	}
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	// we initialize our Users array
	var service ConfigService

	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'users' which we defined above
	json.Unmarshal(byteValue, &service)

	return &service
}

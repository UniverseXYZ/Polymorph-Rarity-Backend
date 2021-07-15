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

// type ComboConfigService struct {
// 	AmishComboCount             int `json:"amishcombocount"`
// 	AstronautComboCount         int `json:"astronautcombocount"`
// 	NinjaComboCount             int `json:"ninjacombocount"`
// 	ClownComboCount             int `json:"clowncombocount"`
// 	ChemicalComboCount          int `json:"chemicalcombocount"`
// 	SamuraiComboCount           int `json:"samuraicombocount"`
// 	RainbowComboCount           int `json:"rainbowcombocount"`
// 	MarineComboCount            int `json:"marinecombocount"`
// 	ZombieRagsComboCount        int `json:"zombieragscombocount"`
// 	HockeyComboCount            int `json:"hockeycombocount"`
// 	SushiChefComboCount         int `json:"sushichefcombocount"`
// 	TaekwondoComboCount         int `json:"taekwondocombocount"`
// 	TennisComboCount            int `json:"tenniscombocount"`
// 	BasketballComboCount        int `json:"basketballcombocount"`
// 	OldFootballStarComboCount   int `json:"oldfootballstarcombocount"`
// 	YoungFootballStarComboCount int `json:"youngfootballstarcombocount"`
// 	StripedSoccerComboCount     int `json:"stripedsoccercombocount"`
// 	SpartanComboCount           int `json:"spartancombocount"`
// 	KnightComboCount            int `json:"knightcombocount"`
// 	GoldenSuitComboCount        int `json:"goldensuitcombocount"`
// 	TuxedoComboCount            int `json:"tuxedocombocount"`
// 	PlaidSuitComboCount         int `json:"plaidsuitcombocount"`
// 	BlackSuitComboCount         int `json:"blacksuitcombocount"`
// 	BrownSuitComboCount         int `json:"brownsuitcombocount"`
// 	GreySuitComboCount          int `json:"greysuitcombocount"`
// 	GolfComboCount              int `json:"golfcombocount"`
// 	SoccerArgentinaComboCount   int `json:"soccerargentinacombocount"`
// 	SoccerBrazilComboCount      int `json:"soccerbrazilcombocount"`
// 	NakedComboCount             int `json:"nakedcombocount"`
// 	StonerComboCount            int `json:"stonercombocount"`
// 	PartyDegenComboCount        int `json:"partydegencombocount"`
// }

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

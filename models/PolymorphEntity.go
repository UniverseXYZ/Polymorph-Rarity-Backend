package models

type PolymorphEntity struct {
	TokenId               string   `json:"tokenid"`
	Rank                  int      `json:"rank"`
	CurrentGene           string   `json:"currentgene"`
	Headwear              string   `json:"headwear"`
	Eyewear               string   `json:"eyewear"`
	Torso                 string   `json:"torso"`
	Pants                 string   `json:"pants"`
	Footwear              string   `json:"footwear"`
	LeftHand              string   `json:"lefthand"`
	RightHand             string   `json:"righthand"`
	Character             string   `json:"character"`
	Background            string   `json:"background"`
	RarityScore           float64  `json:"rarityscore"`
	IsVirgin              bool     `json:"isvirgin"`
	ColorMismatches       int      `json:"colormismatches"`
	MainSetName           string   `json:"mainsetname"`
	MainMatchingTraits    []string `json:"mainmatchingtraits"`
	SecSetName            string   `json:"secsetname"`
	SecMatchingTraits     []string `json:"secmatchingtraits"`
	Owner                 string   `json:"owner"`
	HasCompletedSet       bool     `json:"hascompletedset"`
	HandsScaler           float64  `json:"handsscaler"`
	NoColorMismatchScaler float64  `json:"nocolormismatchscaler"`
	ColorMismatchScaler   float64  `json:"colormismatchscaler"`
	DegenScaler           float64  `json:"degenscaler"`
	VirginScaler          float64  `json:"virginscaler"`
	BaseRarity            float64  `json:"baserarity"`
}

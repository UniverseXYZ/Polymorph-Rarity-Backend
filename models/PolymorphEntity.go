package models

type PolymorphEntity struct {
	TokenId               string   `json:"tokenid,omitempty"`
	Rank                  int      `json:"rank,omitempty"`
	CurrentGene           string   `json:"currentgene,omitempty"`
	Headwear              string   `json:"headwear,omitempty"`
	Eyewear               string   `json:"eyewear,omitempty"`
	Torso                 string   `json:"torso,omitempty"`
	Pants                 string   `json:"pants,omitempty"`
	Footwear              string   `json:"footwear,omitempty"`
	LeftHand              string   `json:"lefthand,omitempty"`
	RightHand             string   `json:"righthand,omitempty"`
	Character             string   `json:"character,omitempty"`
	Background            string   `json:"background,omitempty"`
	RarityScore           float64  `json:"rarityscore,omitempty"`
	IsVirgin              bool     `json:"isvirgin,omitempty"`
	ColorMismatches       int      `json:"colormismatches,omitempty"`
	MainSetName           string   `json:"mainsetname,omitempty"`
	MainMatchingTraits    []string `json:"mainmatchingtraits,omitempty"`
	SecSetName            string   `json:"secsetname,omitempty"`
	SecMatchingTraits     []string `json:"secmatchingtraits,omitempty"`
	Owner                 string   `json:"owner,omitempty"`
	HasCompletedSet       bool     `json:"hascompletedset,omitempty"`
	HandsScaler           float64  `json:"handsscaler,omitempty"`
	NoColorMismatchScaler float64  `json:"nocolormismatchscaler,omitempty"`
	ColorMismatchScaler   float64  `json:"colormismatchscaler,omitempty"`
	DegenScaler           float64  `json:"degenscaler,omitempty"`
	VirginScaler          float64  `json:"virginscaler,omitempty"`
	BaseRarity            float64  `json:"baserarity,omitempty"`
}

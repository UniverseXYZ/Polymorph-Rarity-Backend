package models

type PolymorphEntity struct {
	TokenId         string   `json:"tokenid,omitempty"`
	Rank            int      `json:"rank,omitempty"`
	CurrentGene     string   `json:"currentgene,omitempty"`
	Headwear        string   `json:"headwear,omitempty"`
	Eyewear         string   `json:"eyewear,omitempty"`
	Torso           string   `json:"torso,omitempty"`
	Pants           string   `json:"pants,omitempty"`
	Footwear        string   `json:"footwear,omitempty"`
	LeftHand        string   `json:"lefthand,omitempty"`
	RightHand       string   `json:"righthand,omitempty"`
	Character       string   `json:"character,omitempty"`
	Background      string   `json:"background,omitempty"`
	RarityScore     int      `json:"rarityscore,omitempty"`
	IsVirgin        bool     `json:"isvirgin,omitempty"`
	ColorMismatches int      `json:"colormismatches,omitempty"`
	MatchingTraits  []string `json:"matchingtraits,omitempty"`
}

package models

type PolymorphEntity struct {
	TokenId     string `json:"tokenid,omitempty"`
	Gene        string `json:"gene,omitempty"`
	RarityScore int    `json:"rarityscore,omitempty"`
	IsVirgin    bool   `json:"isvirgin,omitempty"`
}

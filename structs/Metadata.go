package structs

type Metadata struct {
	Description string      `json:"description"`
	Name        string      `json:"name"`
	Image       string      `json:"image"`
	Image3D     string      `json:"image3d"`
	Attributes  []Attribute `json:"attributes"`
	ExternalUrl string      `json:"external_url"`
}

type Attribute struct {
	TraitType string   `json:"trait_type"`
	Value     string   `json:"value"`
	Sets      []string `json:"sets"`
}

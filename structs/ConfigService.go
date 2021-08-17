package structs

type AttributeSet struct {
	Name string   `json:"name"`
	Sets []string `json:"sets"`
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

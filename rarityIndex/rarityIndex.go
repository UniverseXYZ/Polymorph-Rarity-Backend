package rarityIndex

import (
	"log"
	"math"
	"rarity-backend/metadata"
	"strconv"
	"strings"
)

const (
	NO_COLOR_MISMATCH_SCALER = 3
	COLOR_MISMATCH_SCALER    = 1.5
	DEGEN_SCALER             = 0.5
	VIRGIN_SCALER            = 1.5
	MATCHING_HANDS_SCALER    = 1.25
	MISMATCH_PENALTY         = 0.5
)

type SetWithColors struct {
	Name           string
	Colors         []string
	TraitsNumber   float64
	NonColorTraits float64
}

func CalulateRarityScore(attributes []metadata.Attribute, isVirgin bool) int {
	var sets []string
	var leftHand metadata.Attribute
	var rightHand metadata.Attribute
	var virginScaler float64 = 1

	for _, attr := range attributes {
		sets = append(sets, attr.Sets...)
		if attr.TraitType == "Right Hand" {
			rightHand = attr
		} else if attr.TraitType == "Left Hand" {
			leftHand = attr
		}
	}

	hasCompletedSet, setName, matchingTraits := calculateCompleteSets(sets)
	colorMismatches := getColorMismatches(attributes, setName)
	correctHandsScaler := getFullSetHandsScaler(hasCompletedSet, setName, leftHand, rightHand)
	noColorMismatchScaler, colorMismatchScaler, degenScaler, virginScaler := getScalers(hasCompletedSet, setName, colorMismatches, isVirgin)

	baseRarity := math.Pow(2, (matchingTraits - (MISMATCH_PENALTY * colorMismatches)))
	// TODO: Ask ryan about scalers sum/multiply
	totalScalars := virginScaler * correctHandsScaler * noColorMismatchScaler * colorMismatchScaler * degenScaler
	scaledRarity := int(math.Ceil(baseRarity * totalScalars))
	log.Println("Rarity index: " + strconv.Itoa(scaledRarity))
	return scaledRarity
}

func getScalers(hasCompletedSet bool, setName string, colorMismatches float64, isVirgin bool) (float64, float64, float64, float64) {
	var noColorMismatchScaler, colorMismatchScaler, degenScaler, virginScaler float64 = 1, 1, 1, 1

	if hasCompletedSet && colorMismatches == 0 {
		noColorMismatchScaler = NO_COLOR_MISMATCH_SCALER
	} else if hasCompletedSet && colorMismatches != 0 {
		colorMismatchScaler = COLOR_MISMATCH_SCALER
	}

	if setName == "Party Degen" {
		degenScaler = DEGEN_SCALER
	}

	if isVirgin {
		virginScaler = VIRGIN_SCALER
	}

	return noColorMismatchScaler, colorMismatchScaler, degenScaler, virginScaler
}

func getColorMismatches(attributes []metadata.Attribute, longestSet string) float64 {
	footbalSetWithColors := SetWithColors{Name: "Football Star", Colors: []string{"Red", "White", "Yellow"}}
	spartanSetWithColors := SetWithColors{Name: "Spartan", Colors: []string{"Platinum", "Silver", "Gold", "Brown"}}
	knightSetWithColors := SetWithColors{Name: "Knight", Colors: []string{"Silver", "Golden"}}

	var correctSet SetWithColors
	if strings.Contains(longestSet, footbalSetWithColors.Name) {
		correctSet = footbalSetWithColors
	} else if strings.Contains(longestSet, spartanSetWithColors.Name) {
		correctSet = spartanSetWithColors
	} else if strings.Contains(longestSet, knightSetWithColors.Name) {
		correctSet = knightSetWithColors
	} else {
		// Set is without colors
		return 0
	}
	colorMap := make(map[string]float64)
	var totalColorsOccurances, primaryColorOccurances float64

	for _, attr := range attributes {
		for _, color := range correctSet.Colors {
			if strings.Contains(attr.Value, color) {
				totalColorsOccurances++
				colorMap[color]++
				break
			}
		}
	}

	for _, v := range colorMap {
		if primaryColorOccurances < v {
			primaryColorOccurances = v
		}
	}

	colorMismatches := totalColorsOccurances - primaryColorOccurances

	return colorMismatches
}

func getFullSetHandsScaler(hasCompletedSet bool, completedSetName string,
	leftHandAttr metadata.Attribute, rightHandAttr metadata.Attribute) float64 {
	if !hasCompletedSet {
		return 1
	}

	handsMap := map[string][]string{
		"Amish":               {"Amish Pitch Fork"},
		"Astronaut":           {"Naked"},
		"Ninja":               {"Katana", "Bow"},
		"Clown":               {"Naked"},
		"Chemical":            {"Naked"},
		"Samurai":             {"Katana", "Bow"},
		"Rainbow":             {"Naked"},
		"Marine":              {"Grenade", "Big Gun", "Black Gun"},
		"Zombie Rags":         {"Naked"},
		"Hockey":              {"Hockey Stick"},
		"Sushi Chef":          {"Sushi Knife"},
		"Taekwondo":           {"Naked"},
		"Tennis":              {"Tennis Racket"},
		"Old Football Star":   {"American Football"},
		"Young Football Star": {"American Football"},
		"Striped Soccer":      {"Naked"},
		"Spartan":             {"Silver Spartan Sword", "Golden Spartan Sword", "Platinum Spartan Sword", "Shield", "Bow & Arrow"},
		"Basketball":          {"Basketball"},
		"Knight":              {"Sword", "Shield", "Bow & Arrow"},
		"Tuxedo":              {"Big Gun"},
		"Plaid Suit":          {"Naked"},
		"Golden Suit":         {"Golden Gun"},
		"Black Suit":          {"Black Gun"},
		"Brown Suit":          {"Naked"},
		"Grey Suit":           {"Naked"},
		"Golf":                {"Golf Club"},
		"Soccer Argentina":    {"Naked"},
		"Soccer Brazil":       {"Naked"},
		"Naked":               {"Naked"},
		"Stoner":              {"Bong"},
		"Party Degen":         {"Banana", "Bong", "Beer", "Blue Degen Sword", "Double Degen SwordBlue", "Double Degen SwordRed", "Double Degen SwordYellow", "Green Degen Sword", "Purple Degen Sword", "Red Degen Sword"},
	}

	var matchingHandsCount int
	for _, curr := range handsMap[completedSetName] {
		if curr == leftHandAttr.Value || curr == rightHandAttr.Value {
			matchingHandsCount++
		}
	}
	// Here we can easily add another if statement if we implement different scalars for one/two matching hands
	if matchingHandsCount != 0 {
		return MATCHING_HANDS_SCALER
	}
	return 1
}

func calculateCompleteSets(sets []string) (bool, string, float64) {
	var hasCompletedSet bool
	var longestSet int
	var longestSetName string

	setMap := make(map[string]int)

	combosMap := map[string]int{
		"Amish":               5,
		"Astronaut":           4,
		"Ninja":               6,
		"Clown":               4,
		"Chemical":            4,
		"Samurai":             5,
		"Rainbow":             3,
		"Marine":              6,
		"Zombie Rags":         2,
		"Hockey":              5,
		"Sushi Chef":          5,
		"Taekwondo":           2,
		"Tennis":              5,
		"Striped Soccer":      3,
		"Basketball":          4,
		"Tuxedo":              4,
		"Old Football Star":   5,
		"Young Football Star": 5,
		"Spartan":             6,
		"Knight":              6,
		"Golden Suit":         4,
		"Plaid Suit":          5,
		"Black Suit":          6,
		"Brown Suit":          5,
		"Grey Suit":           5,
		"Golf":                5,
		"Soccer Argentina":    3,
		"Soccer Brazil":       3,
		"Naked":               7,
		"Stoner":              3,
		"Party Degen":         7,
	}
	for _, set := range sets {
		setMap[set]++
		if setMap[set] == combosMap[set] {
			hasCompletedSet = true
		}
	}

	for k, v := range setMap {
		if longestSet < v {
			longestSetName, longestSet = k, v
		}
	}

	return hasCompletedSet, longestSetName, float64(longestSet)
}

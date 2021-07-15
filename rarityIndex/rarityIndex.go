package rarityIndex

import (
	"log"
	"math"
	"rarity-backend/metadata"
	"rarity-backend/config"
	"strconv"
	"strings"
)

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

	baseRarity := math.Pow(2, (matchingTraits - (config.MISMATCH_PENALTY * colorMismatches)))
	// TODO: Ask ryan about scalers sum/multiply
	totalScalars := virginScaler * correctHandsScaler * noColorMismatchScaler * colorMismatchScaler * degenScaler
	scaledRarity := int(math.Ceil(baseRarity * totalScalars))
	log.Println("Rarity index: " + strconv.Itoa(scaledRarity))
	return scaledRarity
}

func getScalers(hasCompletedSet bool, setName string, colorMismatches float64, isVirgin bool) (float64, float64, float64, float64) {
	var noColorMismatchScaler, colorMismatchScaler, degenScaler, virginScaler float64 = 1, 1, 1, 1

	if hasCompletedSet && colorMismatches == 0 {
		noColorMismatchScaler = config.NO_COLOR_MISMATCH_SCALER
	} else if hasCompletedSet && colorMismatches != 0 {
		colorMismatchScaler = config.COLOR_MISMATCH_SCALER
	}

	if setName == "Party Degen" {
		degenScaler = config.DEGEN_SCALER
	}

	if isVirgin {
		virginScaler = config.VIRGIN_SCALER
	}

	return noColorMismatchScaler, colorMismatchScaler, degenScaler, virginScaler
}

func getColorMismatches(attributes []metadata.Attribute, longestSet string) float64 {
	footbalSetWithColors := config.SetWithColors{Name: "Football Star", Colors: []string{"Red", "White", "Yellow"}}
	spartanSetWithColors := config.SetWithColors{Name: "Spartan", Colors: []string{"Platinum", "Silver", "Gold", "Brown"}}
	knightSetWithColors := config.SetWithColors{Name: "Knight", Colors: []string{"Silver", "Golden"}}

	var correctSet config.SetWithColors
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

	var matchingHandsCount int
	for _, curr := range config.HandsMap[completedSetName] {
		if curr == leftHandAttr.Value || curr == rightHandAttr.Value {
			matchingHandsCount++
		}
	}
	// Here we can easily add another if statement if we implement different scalars for one/two matching hands
	if matchingHandsCount != 0 {
		return config.MATCHING_HANDS_SCALER
	}
	return 1
}

func calculateCompleteSets(sets []string) (bool, string, float64) {
	var hasCompletedSet bool
	var longestSet int
	var longestSetName string

	setMap := make(map[string]int)

	for _, set := range sets {
		setMap[set]++
		if setMap[set] == config.CombosMap[set] {
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

package rarityIndex

import (
	"fmt"
	"log"
	"math"
	"rarity-backend/config"
	"rarity-backend/metadata"
	"rarity-backend/rarityTypes"
	"strings"
)

func CalulateRarityScore(attributes []metadata.Attribute, isVirgin bool) rarityTypes.RarityResult {
	var leftHand, rightHand metadata.Attribute
	var virginScaler float64 = 1
	var rarityAttributes []metadata.Attribute

	for _, attr := range attributes {
		switch attr.TraitType {
		case "Background", "Character":
		case "Right Hand":
			rightHand = attr
		case "Left Hand":
			leftHand = attr
		default:
			rarityAttributes = append(rarityAttributes, attr)
		}
	}

	hasCompletedSet, setName, matchingTraitsCount, mainMatchingTraits, secSetname, secMatchingTraits := calculateCompleteSets(rarityAttributes)
	colorMismatches := getColorMismatches(attributes, setName)
	correctHandsScaler := getFullSetHandsScaler(hasCompletedSet, setName, leftHand, rightHand)
	noColorMismatchScaler, colorMismatchScaler, degenScaler, virginScaler := getScalers(hasCompletedSet, setName, colorMismatches, isVirgin)

	baseRarity := math.Pow(2, (matchingTraitsCount - (config.MISMATCH_PENALTY * colorMismatches)))
	// (No color mismatches scaler/Color mismatches scaler) * Hands scaler / Degen scaler  ) + Virgin scaler)
	totalScalars := (noColorMismatchScaler * colorMismatchScaler * correctHandsScaler * degenScaler)
	scaledRarity := (math.Round((baseRarity * totalScalars * virginScaler * 100)) / 100)
	log.Println("Rarity index: " + fmt.Sprintf("%f", (scaledRarity)))

	return rarityTypes.RarityResult{
		HasCompletedSet:       hasCompletedSet,
		MainSetName:           setName,
		MainMatchingTraits:    mainMatchingTraits,
		SecSetName:            secSetname,
		SecMatchingTraits:     secMatchingTraits,
		ColorMismatches:       colorMismatches,
		HandsScaler:           correctHandsScaler,
		NoColorMismatchScaler: noColorMismatchScaler,
		ColorMismatchScaler:   colorMismatchScaler,
		DegenScaler:           degenScaler,
		VirginScaler:          virginScaler,
		BaseRarity:            baseRarity,
		ScaledRarity:          scaledRarity,
	}
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
	var matchingHandsCount int
	for _, curr := range config.HandsMap[completedSetName] {
		if curr == leftHandAttr.Value || curr == rightHandAttr.Value {
			matchingHandsCount++
		}
	}

	if !hasCompletedSet {
		if matchingHandsCount == 1 {
			return config.NO_SET_ONE_MATCHING_HANDS_SCALER
		} else if matchingHandsCount == 2 && leftHandAttr.Value != rightHandAttr.Value {
			return config.NO_SET_TWO_MATCHING_HANDS_SCALER
		} else if matchingHandsCount == 2 && leftHandAttr.Value == rightHandAttr.Value {
			return config.NO_SET_TWO_SAME_MATCHING_HANDS_SCALER
		}
	} else if hasCompletedSet {
		if matchingHandsCount == 1 {
			return config.HAS_SET_ONE_MATCHING_HANDS_SCALER
		} else if matchingHandsCount == 2 && leftHandAttr.Value != rightHandAttr.Value {
			return config.HAS_SET_TWO_MATCHING_HANDS_SCALER
		} else if matchingHandsCount == 2 && leftHandAttr.Value == rightHandAttr.Value {
			return config.NO_SET_TWO_SAME_MATCHING_HANDS_SCALER
		}
	}
	return 1
}

func calculateCompleteSets(attributes []metadata.Attribute) (bool, string, float64, []string, string, []string) {
	var hasCompletedSet bool
	var mainSet int
	var mainSetName string

	setMap := make(map[string]int)
	setTraitsMap := make(map[string][]string)

	for _, attr := range attributes {
		for _, set := range attr.Sets {
			setMap[set]++
			setTraitsMap[set] = append(setTraitsMap[set], attr.TraitType)
			if setMap[set] == config.CombosMap[set] {
				hasCompletedSet = true
				mainSetName = set
				mainSet = setMap[set]
			}
		}
	}

	if mainSet == 0 {
		for k, v := range setMap {
			if v >= 2 && mainSet < v {
				mainSetName, mainSet = k, v
			}
		}
	}

	var secondarySetCount int
	var secondarySetName string

	for k, v := range setMap {
		if v >= 2 && secondarySetCount < v && k != mainSetName {
			secondarySetName, secondarySetCount = k, v
		}
	}
	mainMatchingTraits := setTraitsMap[mainSetName]
	secondaryMatchingTraits := setTraitsMap[secondarySetName]

	// It would be bad to have degen as main set while you have secondary set with the same number of traits
	if len(mainMatchingTraits) == len(secondaryMatchingTraits) && mainSetName == "Party Degen" {
		mainSetName, secondarySetName = secondarySetName, mainSetName
		mainMatchingTraits, secondaryMatchingTraits = secondaryMatchingTraits, mainMatchingTraits
	}

	return hasCompletedSet, mainSetName, float64(mainSet), mainMatchingTraits, secondarySetName, secondaryMatchingTraits
}

// TODO: Corn gun and diamond hands are left out

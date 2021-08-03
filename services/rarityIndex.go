package services

import (
	"fmt"
	"log"
	"math"
	"rarity-backend/config"
	"rarity-backend/constants"
	"rarity-backend/helpers"
	"rarity-backend/structs"
	"strings"
)

// CalulateRarityScore is the core function responsible for calcualting the rarity score.
//
// It calculates the rarity score of the polymorph, the different scalers used in the formuala and other rarity related metadata that is tracked and stored in the database.
//
// Configurations can be found in rarityConfig.go
func CalulateRarityScore(attributes []structs.Attribute, isVirgin bool) structs.RarityResult {
	leftHand, rightHand, rarityAttributes := parseAttributes(attributes)

	hasCompletedSet, setName, mainMatchingTraits, secSetname, secMatchingTraits := calculateCompleteSets(rarityAttributes)
	isColoredSet, colorMismatches := getColorMismatches(attributes, setName)
	scalers := getScalers(hasCompletedSet, setName, colorMismatches, isVirgin, isColoredSet)
	handsScaler, handsSetName, matchingHandsCount, mainMatchingTraitsWithHands := getFullSetHandsScaler(mainMatchingTraits, hasCompletedSet, setName, leftHand, rightHand)

	mainSetCount := float64(len(mainMatchingTraits))
	secSetBonus := config.SECONDARY_SET_SCALER * float64(len(secMatchingTraits))
	mismatchPenalty := config.MISMATCH_PENALTY * colorMismatches

	baseRarity := math.Pow(2, mainSetCount-mismatchPenalty+secSetBonus)

	totalScalars := scalers.NoColorMismatchScaler * scalers.ColorMismatchScaler * handsScaler * scalers.DegenScaler * scalers.VirginScaler
	scaledRarity := math.Round((baseRarity * totalScalars * 100)) / 100
	log.Println("Rarity index: " + fmt.Sprintf("%f", (scaledRarity)))

	return structs.RarityResult{
		HasCompletedSet:       hasCompletedSet,
		MainSetName:           setName,
		MainMatchingTraits:    mainMatchingTraitsWithHands,
		SecSetName:            secSetname,
		SecMatchingTraits:     secMatchingTraits,
		ColorMismatches:       int(colorMismatches),
		HandsScaler:           handsScaler,
		HandsSetName:          handsSetName,
		MatchingHands:         matchingHandsCount,
		NoColorMismatchScaler: scalers.NoColorMismatchScaler,
		ColorMismatchScaler:   scalers.ColorMismatchScaler,
		DegenScaler:           scalers.DegenScaler,
		VirginScaler:          scalers.VirginScaler,
		BaseRarity:            baseRarity,
		ScaledRarity:          scaledRarity,
	}
}

//parseAttributes parses the array of attrbutes.
//
//Returns left hand, right hand and the attrbitues without Character, Background attrbiutes as they aren't used in the rarity score formula
func parseAttributes(attributes []structs.Attribute) (structs.Attribute, structs.Attribute, []structs.Attribute) {
	var leftHand, rightHand structs.Attribute
	var rarityAttributes []structs.Attribute

	for _, attr := range attributes {
		switch attr.TraitType {
		case constants.MorphAttriutes.Background, constants.MorphAttriutes.Character:
		case constants.MorphAttriutes.RightHand:
			rightHand = attr
		case constants.MorphAttriutes.LeftHand:
			leftHand = attr
		default:
			rarityAttributes = append(rarityAttributes, attr)
		}
	}

	return leftHand, rightHand, rarityAttributes
}

// getScalers calculates the eligible scalers for the polymorph
func getScalers(hasCompletedSet bool, setName string, colorMismatches float64, isVirgin bool, isColoredSet bool) structs.Scalers {
	var noColorMismatchScaler, colorMismatchScaler, degenScaler, virginScaler float64 = 1, 1, 1, 1

	if hasCompletedSet && isColoredSet && colorMismatches == 0 {
		noColorMismatchScaler = config.NO_COLOR_MISMATCH_SCALER
	} else if hasCompletedSet && isColoredSet && colorMismatches != 0 {
		colorMismatchScaler = config.COLOR_MISMATCH_SCALER
	}

	if setName == "Party Degen" {
		degenScaler = config.DEGEN_SCALER
	}

	if isVirgin {
		virginScaler = config.VIRGIN_SCALER
	}

	return structs.Scalers{
		ColorMismatchScaler:   colorMismatchScaler,
		NoColorMismatchScaler: noColorMismatchScaler,
		VirginScaler:          virginScaler,
		DegenScaler:           degenScaler,
	}
}

// getColorMismatches calculates determines if the set has colors or not and the number of color mismatches if applicable.
//
// Color sets can be found in rarityConfig.go
func getColorMismatches(attributes []structs.Attribute, longestSet string) (bool, float64) {
	var correctSet structs.ColorSet
	if strings.Contains(longestSet, config.FootbalSetWithColors.Name) {
		correctSet = config.FootbalSetWithColors
	} else if strings.Contains(longestSet, config.SpartanSetWithColors.Name) {
		correctSet = config.SpartanSetWithColors
	} else if strings.Contains(longestSet, config.KnightSetWithColors.Name) {
		correctSet = config.KnightSetWithColors
	} else {
		// Set is without colors
		return false, 0
	}
	colorMap := make(map[string]float64)
	var totalColorsOccurances, primaryColorOccurances float64

	for _, attr := range attributes {
		for _, color := range correctSet.Colors {
			if strings.Contains(attr.Value, color) && helpers.StringInSlice(longestSet, attr.Sets) {
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

	return true, colorMismatches
}

// getFullSetHandsScaler calculates the correct hands scaler based on the state of the set(no, incomplete or completed set)
func getFullSetHandsScaler(mainMatchingTraits []string, hasCompletedSet bool, completedSetName string,
	leftHandAttr structs.Attribute, rightHandAttr structs.Attribute) (float64, string, int, []string) {
	var matchingSetHandsCount int
	for _, handAttribute := range config.HandsMap[completedSetName] {
		if handAttribute == leftHandAttr.Value {
			matchingSetHandsCount++
			mainMatchingTraits = append(mainMatchingTraits, leftHandAttr.TraitType)
		} else if handAttribute == rightHandAttr.Value {
			matchingSetHandsCount++
			mainMatchingTraits = append(mainMatchingTraits, rightHandAttr.TraitType)
		}
	}
	if matchingSetHandsCount == 0 {
		// Check if they are matching set
		handMap := make(map[string]int)
		for _, set := range leftHandAttr.Sets {
			handMap[set]++
		}
		for _, set := range rightHandAttr.Sets {
			handMap[set]++
			if handMap[set] == 2 {
				if leftHandAttr.Value == rightHandAttr.Value {
					return config.NO_SET_TWO_SAME_MATCHING_HANDS_SCALER, set, handMap[set], mainMatchingTraits
				} else {
					return config.NO_SET_TWO_MATCHING_HANDS_SCALER, set, handMap[set], mainMatchingTraits
				}
			}
		}
	} else if !hasCompletedSet {
		if matchingSetHandsCount == 1 {
			return config.INCOMPLETE_SET_ONE_MATCHING_HANDS_SCALER, completedSetName, matchingSetHandsCount, mainMatchingTraits
		}
		if matchingSetHandsCount == 2 && leftHandAttr.Value != rightHandAttr.Value {
			return config.INCOMPLETE_SET_TWO_MATCHING_HANDS_SCALER, completedSetName, matchingSetHandsCount, mainMatchingTraits
		}
		if matchingSetHandsCount == 2 && leftHandAttr.Value == rightHandAttr.Value {
			return config.INCOMPLETE_SET_TWO_SAME_MATCHING_HANDS_SCALER, completedSetName, matchingSetHandsCount, mainMatchingTraits
		}
	} else if hasCompletedSet {
		if matchingSetHandsCount == 1 {
			return config.HAS_SET_ONE_MATCHING_HANDS_SCALER, completedSetName, matchingSetHandsCount, mainMatchingTraits
		}
		if matchingSetHandsCount == 2 && leftHandAttr.Value != rightHandAttr.Value {
			return config.HAS_SET_TWO_MATCHING_HANDS_SCALER, completedSetName, matchingSetHandsCount, mainMatchingTraits
		}
		if matchingSetHandsCount == 2 && leftHandAttr.Value == rightHandAttr.Value {
			return config.INCOMPLETE_SET_TWO_SAME_MATCHING_HANDS_SCALER, completedSetName, matchingSetHandsCount, mainMatchingTraits
		}
	}
	return 1, "", 0, mainMatchingTraits
}

// calculateCompleteSets iterates over polymorph's attributes.
//
// Return if set has been completed, main set name, main set attrbiutes, secondary set name, secondary set attributes
func calculateCompleteSets(attributes []structs.Attribute) (bool, string, []string, string, []string) {
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

	return hasCompletedSet, mainSetName, mainMatchingTraits, secondarySetName, secondaryMatchingTraits
}

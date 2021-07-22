package helpers

import (
	"rarity-backend/metadata"
	"rarity-backend/models"
	"rarity-backend/rarityTypes"
	"sort"

	"github.com/ethereum/go-ethereum/core/types"
)

func CreateMorphEntity(event rarityTypes.PolymorphEvent, attributes []metadata.Attribute, isVirgin bool, rarityResult rarityTypes.RarityResult) models.PolymorphEntity {
	var background, leftHand, rightHand, head, eye, torso, pants, feet, character metadata.Attribute

	for _, attr := range attributes {
		switch attr.TraitType {
		case "Background":
			background = attr
		case "Character":
			character = attr
		case "Right Hand":
			rightHand = attr
		case "Left Hand":
			leftHand = attr
		case "Footwear":
			feet = attr
		case "Pants":
			pants = attr
		case "Torso":
			torso = attr
		case "Eyewear":
			eye = attr
		case "Headwear":
			head = attr
		}
	}

	morphEntity := models.PolymorphEntity{
		TokenId:               event.MorphId.String(),
		Rank:                  0,
		CurrentGene:           event.NewGene.String(),
		Headwear:              head.Value,
		Eyewear:               eye.Value,
		Torso:                 torso.Value,
		Pants:                 pants.Value,
		Footwear:              feet.Value,
		LeftHand:              leftHand.Value,
		RightHand:             rightHand.Value,
		Character:             character.Value,
		Background:            background.Value,
		RarityScore:           rarityResult.ScaledRarity,
		IsVirgin:              isVirgin,
		ColorMismatches:       int(rarityResult.ColorMismatches),
		MainSetName:           rarityResult.MainSetName,
		MainMatchingTraits:    rarityResult.MainMatchingTraits,
		SecSetName:            rarityResult.SecSetName,
		SecMatchingTraits:     rarityResult.SecMatchingTraits,
		HasCompletedSet:       rarityResult.HasCompletedSet,
		HandsScaler:           rarityResult.ColorMismatches,
		NoColorMismatchScaler: rarityResult.NoColorMismatchScaler,
		ColorMismatchScaler:   rarityResult.ColorMismatches,
		DegenScaler:           rarityResult.DegenScaler,
		VirginScaler:          rarityResult.VirginScaler,
		BaseRarity:            rarityResult.BaseRarity,
	}
	return morphEntity
}

func SortMorphEvents(eventLogs []types.Log) {
	sort.Slice(eventLogs, func(i, j int) bool {
		curr := eventLogs[i]
		prev := eventLogs[j]

		if curr.BlockNumber < prev.BlockNumber {
			return true
		}

		if curr.BlockNumber > prev.BlockNumber {
			return false
		}

		if curr.TxIndex < prev.TxIndex {
			return true
		}

		if curr.TxIndex > prev.TxIndex {
			return false
		}

		if curr.Index < prev.Index {
			return true
		}

		if curr.Index > prev.Index {
			return false
		}
		return true
	})

}

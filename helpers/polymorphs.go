package helpers

import (
	"rarity-backend/constants"
	"rarity-backend/models"
	"rarity-backend/structs"
	"sort"

	"github.com/ethereum/go-ethereum/core/types"
)

func CreateMorphEntity(event structs.PolymorphEvent, attributes []structs.Attribute, isVirgin bool, rarityResult structs.RarityResult) models.PolymorphEntity {
	var background, leftHand, rightHand, head, eye, torso, pants, feet, character structs.Attribute

	for _, attr := range attributes {
		switch attr.TraitType {
		case constants.MorphAttriutes.Background:
			background = attr
		case constants.MorphAttriutes.Character:
			character = attr
		case constants.MorphAttriutes.RightHand:
			rightHand = attr
		case constants.MorphAttriutes.LeftHand:
			leftHand = attr
		case constants.MorphAttriutes.Footwear:
			feet = attr
		case constants.MorphAttriutes.Pants:
			pants = attr
		case constants.MorphAttriutes.Torso:
			torso = attr
		case constants.MorphAttriutes.Eyewear:
			eye = attr
		case constants.MorphAttriutes.Headwear:
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
		ColorMismatches:       rarityResult.ColorMismatches,
		MainSetName:           rarityResult.MainSetName,
		MainMatchingTraits:    rarityResult.MainMatchingTraits,
		SecSetName:            rarityResult.SecSetName,
		SecMatchingTraits:     rarityResult.SecMatchingTraits,
		HasCompletedSet:       rarityResult.HasCompletedSet,
		HandsScaler:           rarityResult.HandsScaler,
		HandsSetName:          rarityResult.HandsSetName,
		MatchingHands:         rarityResult.MatchingHands,
		NoColorMismatchScaler: rarityResult.NoColorMismatchScaler,
		ColorMismatchScaler:   rarityResult.ColorMismatchScaler,
		DegenScaler:           rarityResult.DegenScaler,
		VirginScaler:          rarityResult.VirginScaler,
		BaseRarity:            rarityResult.BaseRarity,
	}
	if len(morphEntity.SecMatchingTraits) == 0 {
		morphEntity.SecMatchingTraits = []string{}
	}
	if len(morphEntity.MainMatchingTraits) == 0 {
		morphEntity.MainMatchingTraits = []string{}
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

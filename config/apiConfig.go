package config

import "rarity-backend/constants"

var SEARCH_QUERY_FIELDS []string = []string{
	constants.MorphFieldNames.TokenId,
	constants.MorphFieldNames.Rank,
	constants.MorphFieldNames.RarityScore,
	constants.MorphFieldNames.Headwear,
	constants.MorphFieldNames.Eyewear,
	constants.MorphFieldNames.Torso,
	constants.MorphFieldNames.Pants,
	constants.MorphFieldNames.Footwear,
	constants.MorphFieldNames.LeftHand,
	constants.MorphFieldNames.RightHand,
	constants.MorphFieldNames.Character,
	constants.MorphFieldNames.MainSetName,
	constants.MorphFieldNames.SecSetName,
}

var NO_PROJECTION_FIELDS []string = []string{
	constants.MorphFieldNames.ObjId,
	constants.MorphFieldNames.OldGenes,
	constants.MorphFieldNames.Character,
	constants.MorphFieldNames.Background,
	constants.MorphFieldNames.Headwear,
	constants.MorphFieldNames.Eyewear,
	constants.MorphFieldNames.Torso,
	constants.MorphFieldNames.Pants,
	constants.MorphFieldNames.Footwear,
	constants.MorphFieldNames.LeftHand,
	constants.MorphFieldNames.RightHand,
}

const RESULTS_LIMIT int64 = 100
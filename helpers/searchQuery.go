package helpers

import (
	"rarity-backend/config"
	"rarity-backend/constants"
	"strconv"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ParseSearchQueryString(search string) bson.M {
	queries := bson.A{}

	for _, field := range config.SEARCH_QUERY_FIELDS {
		pattern := ""
		switch field {
		case constants.MorphFieldNames.TokenId:
			pattern = getExactTokenPattern(search)
			regex := primitive.Regex{Pattern: pattern, Options: "i"}
			regexFilter := bson.M{"$regex": regex}
			queries = append(queries, bson.M{field: regexFilter})
		case constants.MorphFieldNames.Rank,
			constants.MorphFieldNames.RarityScore:
			parsed, err := strconv.Atoi(search)
			if err == nil {
				queries = append(queries, bson.M{field: parsed})
			}
		default:
			pattern = search
			regex := primitive.Regex{Pattern: pattern, Options: "i"}
			regexFilter := bson.M{"$regex": regex}
			queries = append(queries, bson.M{field: regexFilter})
		}
	}
	orQuery := bson.M{"$or": queries}
	return orQuery
}

func getExactTokenPattern(number string) string {
	return "(^|\\D)" + number + "(?!\\d)"
}

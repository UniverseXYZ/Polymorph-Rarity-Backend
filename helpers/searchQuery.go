package helpers

import (
	"rarity-backend/config"
	"strconv"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ParseSearchQueryString(search string) bson.M {
	queries := bson.A{}
	// searchQueryFields := []string{"rank", "tokenid", "mainsetname", "secsetname"}

	for _, field := range config.SEARCH_QUERY_FIELDS {
		pattern := ""
		switch field {
		case "tokenid":
			pattern = getExactTokenPattern(search)
			regex := primitive.Regex{Pattern: pattern, Options: "i"}
			regexFilter := bson.M{"$regex": regex}
			queries = append(queries, bson.M{field: regexFilter})
		case "rank":
			rank, err := strconv.Atoi(search)
			if err == nil {
				queries = append(queries, bson.M{field: rank})
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

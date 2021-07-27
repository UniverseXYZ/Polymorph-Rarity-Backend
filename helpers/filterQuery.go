package helpers

import (
	"log"
	"strconv"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
)

const (
	paramSeparator = ";"
	expSeparator   = "_"
)

type Expression struct {
	Field     string
	Operator  string
	Value     string
	Join      string
	Operator2 string
	Value2    string
}

func ParseFilterQueryString(filter string) bson.M {
	// rarityscore_gte_10.2_and_lte_12.4;mainsetname_eq_Spartan;isvirgin_eq_false
	expressions := strings.Split(filter, paramSeparator)
	expArray := []Expression{}

	for _, expression := range expressions {
		exParts := strings.Split(expression, expSeparator)
		if len(exParts) == 3 {
			field, operator, value := strings.ToLower(exParts[0]), exParts[1], exParts[2]

			expArray = append(expArray, Expression{
				Field:    field,
				Operator: operator,
				Value:    value,
			})
		} else if len(exParts) == 6 {
			field, operator, value, join, operator2, value2 := strings.ToLower(exParts[0]), exParts[1], exParts[2], exParts[3], exParts[4], exParts[5]

			expArray = append(expArray, Expression{
				Field:     field,
				Operator:  operator,
				Value:     value,
				Join:      join,
				Operator2: operator2,
				Value2:    value2,
			})

		}
	}

	filters := buildFilter(expArray)
	return filters
}

func buildFilter(expressions []Expression) bson.M {
	filter := bson.M{}
	for _, exp := range expressions {
		switch exp.Join {
		case "":
			switch exp.Operator {
			case "eq":
				currBson := createEqBson(exp.Field, exp.Value)
				for k, v := range currBson {
					filter[k] = v
				}
			case "lt", "lte", "gt", "gte":
				currBson := createCompareBson(exp.Field, exp.Operator, exp.Value)
				for k, v := range currBson {
					filter[k] = v
				}
			}
		case "and", "or":
			var bson1, bson2 bson.M
			bson1 = createCompareBson(exp.Field, exp.Operator, exp.Value)
			bson2 = createCompareBson(exp.Field, exp.Operator2, exp.Value2)

			aBson := bson.A{bson1, bson2}
			filter["$"+exp.Join] = aBson
		}
	}

	return filter
}

func createAndEqBson(field string, value string, existingFilter interface{}) bson.A {
	bson1 := bson.M{}
	bson2 := bson.M{}
	if value == "true" || value == "false" {
		boolValue, err := strconv.ParseBool(value)
		if err != nil {
			log.Println(err)
		} else {
			bson1[field] = boolValue
		}
	} else {
		bson1[field] = value
	}
	bson2[field] = existingFilter
	aBson := bson.A{bson1, bson2}
	return aBson
}

func createEqBson(field string, value string) bson.M {
	returnBson := bson.M{}
	if value == "true" || value == "false" {
		boolValue, err := strconv.ParseBool(value)
		if err != nil {
			log.Println(err)
		} else {
			returnBson[field] = boolValue
		}
	} else {
		returnBson[field] = value
	}
	return returnBson
}

func createCompareBson(field string, operator string, value string) bson.M {
	returnBson := bson.M{}

	floatValue, err := strconv.ParseFloat(value, 64)
	if err != nil {
		log.Println(err)
	} else {
		nestedBson := bson.M{"$" + operator: floatValue}
		returnBson[field] = nestedBson
	}
	return returnBson
}

func createAndCompareBson(field string, operator string, value string, existingFilter interface{}) bson.A {
	bson1 := bson.M{}
	bson2 := bson.M{}
	floatValue, err := strconv.ParseFloat(value, 64)
	if err != nil {
		log.Println(err)
	} else {
		nestedBson := bson.M{"$" + operator: floatValue}
		bson1[field] = nestedBson
	}
	bson2[field] = existingFilter
	aBson := bson.A{bson1, bson2}
	return aBson
}

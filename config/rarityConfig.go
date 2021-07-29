package config

import "rarity-backend/structs"

var SCRAMBLE_COST float32 = 0.01

var NO_COLOR_MISMATCH_SCALER float64 = 3
var COLOR_MISMATCH_SCALER float64 = 1.5
var DEGEN_SCALER float64 = 0.5
var VIRGIN_SCALER float64 = 1.5
var MISMATCH_PENALTY float64 = 0.5
var SECONDARY_SET_SCALER float64 = 0.5

var NO_SET_TWO_MATCHING_HANDS_SCALER float64 = 1.1
var NO_SET_TWO_SAME_MATCHING_HANDS_SCALER float64 = 1.2
var INCOMPLETE_SET_ONE_MATCHING_HANDS_SCALER float64 = 1.3
var INCOMPLETE_SET_TWO_MATCHING_HANDS_SCALER float64 = 1.4
var INCOMPLETE_SET_TWO_SAME_MATCHING_HANDS_SCALER float64 = 1.5
var HAS_SET_ONE_MATCHING_HANDS_SCALER float64 = 1.6
var HAS_SET_TWO_MATCHING_HANDS_SCALER float64 = 1.7
var HAS_SET_TWO_SAME_MATCHING_HANDS_SCALER float64 = 1.8

var FootbalSetWithColors = structs.ColorSet{Name: "Football Star", Colors: []string{"Red", "White", "Yellow"}}
var SpartanSetWithColors = structs.ColorSet{Name: "Spartan", Colors: []string{"Platinum", "Silver", "Gold", "Brown"}}
var KnightSetWithColors = structs.ColorSet{Name: "Knight", Colors: []string{"Silver", "Golden"}}

var HandsMap = map[string][]string{
	"Amish":            {"Amish Pitch Fork"},
	"Astronaut":        {"Naked"},
	"Ninja":            {"Katana", "Bow"},
	"Clown":            {"Naked"},
	"Chemical":         {"Naked"},
	"Samurai":          {"Katana", "Bow"},
	"Rainbow":          {"Naked"},
	"Marine":           {"Grenade", "Big Gun", "Black Gun"},
	"Zombie Rags":      {"Naked"},
	"Hockey":           {"Hockey Stick"},
	"Sushi Chef":       {"Sushi Knife"},
	"Taekwondo":        {"Naked"},
	"Tennis":           {"Tennis Racket"},
	"Football Star":    {"American Football"},
	"Striped Soccer":   {"Naked"},
	"Spartan":          {"Silver Spartan Sword", "Golden Spartan Sword", "Platinum Spartan Sword", "Shield", "Bow & Arrow"},
	"Basketball":       {"Basketball"},
	"Knight":           {"Sword", "Shield", "Bow & Arrow"},
	"Tuxedo":           {"Big Gun"},
	"Plaid Suit":       {"Naked"},
	"Golden Suit":      {"Golden Gun"},
	"Black Suit":       {"Black Gun"},
	"Brown Suit":       {"Naked"},
	"Grey Suit":        {"Naked"},
	"Golf":             {"Golf Club"},
	"Soccer Argentina": {"Naked"},
	"Soccer Brazil":    {"Naked"},
	"Naked":            {"Naked"},
	"Stoner":           {"Bong"},
	"Party Degen":      {"Banana", "Bong", "Beer", "Blue Degen Sword", "Double Degen SwordBlue", "Double Degen SwordRed", "Double Degen SwordYellow", "Green Degen Sword", "Purple Degen Sword", "Red Degen Sword"},
}

var CombosMapNew = map[string]int{
	"Zombie Rags":      2,
	"Taekwondo":        2,
	"Hockey":           3,
	"Tennis":           3,
	"Striped Soccer":   3,
	"Basketball":       3,
	"Grey Suit":        3,
	"Soccer Argentina": 3,
	"Soccer Brazil":    3,
	"Stoner":           3,
	"Party Degen":      3,
	"Samurai":          3,
	"Amish":            4,
	"Astronaut":        4,
	"Ninja":            4,
	"Clown":            4,
	"Chemical":         4,
	"Rainbow":          4,
	"Marine":           4,
	"Sushi Chef":       4,
	"Football Star":    4,
	"Spartan":          4,
	"Knight":           4,
	"Plaid Suit":       4,
	"Black Suit":       4,
	"Brown Suit":       4,
	"Golf":             4,
	"Tuxedo":           5,
	"Golden Suit":      5,
	"Naked":            5,
}

var CombosMap = map[string]int{
	"Amish":            4,
	"Astronaut":        4,
	"Ninja":            4,
	"Clown":            4,
	"Chemical":         4,
	"Samurai":          3,
	"Rainbow":          3,
	"Marine":           4,
	"Zombie Rags":      2,
	"Hockey":           4,
	"Sushi Chef":       4,
	"Taekwondo":        2,
	"Tennis":           3,
	"Striped Soccer":   3,
	"Basketball":       3,
	"Tuxedo":           4,
	"Football Star":    4,
	"Spartan":          4,
	"Knight":           4,
	"Golden Suit":      5,
	"Plaid Suit":       4,
	"Black Suit":       4,
	"Brown Suit":       4,
	"Grey Suit":        4,
	"Golf":             4,
	"Soccer Argentina": 3,
	"Soccer Brazil":    3,
	"Naked":            5,
	"Stoner":           1,
	"Party Degen":      5,
}

package config

var NO_COLOR_MISMATCH_SCALER float64 = 3
var COLOR_MISMATCH_SCALER float64 = 1.5
var DEGEN_SCALER float64 = 0.5
var VIRGIN_SCALER float64 = 1.5
var MISMATCH_PENALTY float64 = 0.5
var NO_SET_ONE_MATCHING_HANDS_SCALER float64 = 1.1
var NO_SET_TWO_MATCHING_HANDS_SCALER float64 = 1.2
var NO_SET_TWO_SAME_MATCHING_HANDS_SCALER float64 = 1.3
var HAS_SET_ONE_MATCHING_HANDS_SCALER float64 = 1.4
var HAS_SET_TWO_MATCHING_HANDS_SCALER float64 = 1.5
var HAS_SET_TWO_SAME_MATCHING_HANDS_SCALER float64 = 1.6

var HandsMap = map[string][]string{
	"Amish":               {"Amish Pitch Fork"},
	"Astronaut":           {"Naked"},
	"Ninja":               {"Katana", "Bow"},
	"Clown":               {"Naked"},
	"Chemical":            {"Naked"},
	"Samurai":             {"Katana", "Bow"},
	"Rainbow":             {"Naked"},
	"Marine":              {"Grenade", "Big Gun", "Black Gun"},
	"Zombie Rags":         {"Naked"},
	"Hockey":              {"Hockey Stick"},
	"Sushi Chef":          {"Sushi Knife"},
	"Taekwondo":           {"Naked"},
	"Tennis":              {"Tennis Racket"},
	"Old Football Star":   {"American Football"},
	"Young Football Star": {"American Football"},
	"Striped Soccer":      {"Naked"},
	"Spartan":             {"Silver Spartan Sword", "Golden Spartan Sword", "Platinum Spartan Sword", "Shield", "Bow & Arrow"},
	"Basketball":          {"Basketball"},
	"Knight":              {"Sword", "Shield", "Bow & Arrow"},
	"Tuxedo":              {"Big Gun"},
	"Plaid Suit":          {"Naked"},
	"Golden Suit":         {"Golden Gun"},
	"Black Suit":          {"Black Gun"},
	"Brown Suit":          {"Naked"},
	"Grey Suit":           {"Naked"},
	"Golf":                {"Golf Club"},
	"Soccer Argentina":    {"Naked"},
	"Soccer Brazil":       {"Naked"},
	"Naked":               {"Naked"},
	"Stoner":              {"Bong"},
	"Party Degen":         {"Banana", "Bong", "Beer", "Blue Degen Sword", "Double Degen SwordBlue", "Double Degen SwordRed", "Double Degen SwordYellow", "Green Degen Sword", "Purple Degen Sword", "Red Degen Sword"},
}

var CombosMap = map[string]int{
	"Amish":               5,
	"Astronaut":           4,
	"Ninja":               6,
	"Clown":               4,
	"Chemical":            4,
	"Samurai":             5,
	"Rainbow":             3,
	"Marine":              6,
	"Zombie Rags":         2,
	"Hockey":              5,
	"Sushi Chef":          5,
	"Taekwondo":           2,
	"Tennis":              5,
	"Striped Soccer":      3,
	"Basketball":          4,
	"Tuxedo":              4,
	"Old Football Star":   5,
	"Young Football Star": 5,
	"Spartan":             6,
	"Knight":              6,
	"Golden Suit":         4,
	"Plaid Suit":          5,
	"Black Suit":          6,
	"Brown Suit":          5,
	"Grey Suit":           5,
	"Golf":                5,
	"Soccer Argentina":    3,
	"Soccer Brazil":       3,
	"Naked":               7,
	"Stoner":              3,
	"Party Degen":         7,
}

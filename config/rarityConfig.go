package config

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

var TokenMintedSignature = "0x5f7666687319b40936f33c188908d86aea154abd3f4127b4fa0a3f04f303c7da"
var TokenMorphedSignature = "0x8c0bdd7bca83c4e0c810cbecf44bc544a9dc0b9f265664e31ce0ce85f07a052b"

// 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef - TRANSFER EVENT
// 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925 - APPROVAL EVENT
// 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31 - APPROVAL FOR ALL EVENT
// 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d - ROLE GRANTED
// 0x5f7666687319b40936f33c188908d86aea154abd3f4127b4fa0a3f04f303c7da - TOKEN MINTED

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

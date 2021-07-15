package main

import (
	"log"
	"os"

	"rarity-backend/config"
	"rarity-backend/dlt"
	"rarity-backend/handlers"
	"rarity-backend/services"

	"github.com/gofiber/fiber"
	"github.com/jasonlvhit/gocron"
	"github.com/joho/godotenv"
)

func connectToEthereum() *dlt.EthereumClient {

	nodeURL := os.Getenv("NODE_URL")

	client, err := dlt.NewEthereumClient(nodeURL)

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Successfully connected to ethereum client")

	return client
}

func initResources() (*dlt.EthereumClient, string, *config.ConfigService) {
	// Load env variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file: " + err.Error())
	}

	// Inital step: Recover to be up to date
	ethClient := connectToEthereum()

	contractAddress := os.Getenv("CONTRACT_ADDRESS")
	if contractAddress == "" {
		log.Fatal("Missing contract address")
	}

	configService := config.NewConfigService("./config.json")

	return ethClient, contractAddress, configService
}

func startAPI() {
	// Routine two: API -> Should start after deploy?
	app := fiber.New()
	app.Get("/morphs/", handlers.GetPolymorphs)
	app.Get("/morphs/:id?", handlers.GetPolymorphById)
	log.Fatal(app.Listen(8000))
}

func recoverAndPoll(ethClient *dlt.EthereumClient, contractAddress string, configService *config.ConfigService) {
	// Recover immediately
	services.ProcessBlocks(ethClient, contractAddress, configService)
	// Routine one: Start polling after recovery
	gocron.Every(15).Second().Do(services.ProcessBlocks, ethClient, contractAddress, configService)
	<-gocron.Start()
}

func main() {
	ethClient, contractAddress, configService := initResources()
	go recoverAndPoll(ethClient, contractAddress, configService)
	startAPI()
}

// func main() {
// // Rarity: 48
// mismatchedSpartanSet := []metadata.Attribute{
// 	{
// 		TraitType: "Eyewear",
// 		Value:     "No Eyewear",
// 		Sets:      []string{"Naked"},
// 	},
// 	{
// 		TraitType: "Footwear",
// 		Value:     "Platinum Spartan Sandals",
// 		Sets:      []string{"Spartan"},
// 	},
// 	{
// 		TraitType: "Headwear",
// 		Value:     "Golden Spartan Helmet",
// 		Sets:      []string{"Spartan"},
// 	},
// 	{
// 		TraitType: "Torso",
// 		Value:     "Silver Spartan Armor",
// 		Sets:      []string{"Spartan"},
// 	},
// 	{
// 		TraitType: "Pants",
// 		Value:     "Spartan Pants",
// 		Sets:      []string{"Spartan"},
// 	},
// 	{
// 		TraitType: "Left Hand",
// 		Value:     "Golden Spartan Sword",
// 		Sets:      []string{"Spartan"},
// 	},
// 	{
// 		TraitType: "Right Hand",
// 		Value:     "Bow & Arrow",
// 		Sets:      []string{"Ninja", "Samurai", "Spartan", "Knight"},
// 	},
// }
// rarityIndex.CalulateRarityScore(mismatchedSpartanSet)

// //Rarity: 240
// matchedKnightSet := []metadata.Attribute{
// 	{
// 		TraitType: "Eyewear",
// 		Value:     "Golden Knight Helmet",
// 		Sets:      []string{"Knight"},
// 	},
// 	{
// 		TraitType: "Footwear",
// 		Value:     "Golden Knight Boots",
// 		Sets:      []string{"Knight"},
// 	},
// 	{
// 		TraitType: "Headwear",
// 		Value:     "Golden Spartan Helmet",
// 		Sets:      []string{"Knight"},
// 	},
// 	{
// 		TraitType: "Torso",
// 		Value:     "Golden Armor",
// 		Sets:      []string{"Knight"},
// 	},
// 	{
// 		TraitType: "Pants",
// 		Value:     "Golden Grieves",
// 		Sets:      []string{"Knight"},
// 	},
// 	{
// 		TraitType: "Left Hand",
// 		Value:     "Shield",
// 		Sets:      []string{"Knight"},
// 	},
// 	{
// 		TraitType: "Right Hand",
// 		Value:     "Bong",
// 		Sets:      []string{"Party Degen"},
// 	},
// }

// rarityIndex.CalulateRarityScore(matchedKnightSet)

// //Rarity: 85
// mismatchedKnightSet := []metadata.Attribute{
// 	{
// 		TraitType: "Eyewear",
// 		Value:     "Golden Knight Helmet",
// 		Sets:      []string{"Knight"},
// 	},
// 	{
// 		TraitType: "Footwear",
// 		Value:     "Silver Knight Boots",
// 		Sets:      []string{"Knight"},
// 	},
// 	{
// 		TraitType: "Headwear",
// 		Value:     "Golden Spartan Helmet",
// 		Sets:      []string{"Knight"},
// 	},
// 	{
// 		TraitType: "Torso",
// 		Value:     "Golden Armor",
// 		Sets:      []string{"Knight"},
// 	},
// 	{
// 		TraitType: "Pants",
// 		Value:     "Golden Grieves",
// 		Sets:      []string{"Knight"},
// 	},
// 	{
// 		TraitType: "Left Hand",
// 		Value:     "Shield",
// 		Sets:      []string{"Knight"},
// 	},
// 	{
// 		TraitType: "Right Hand",
// 		Value:     "Bong",
// 		Sets:      []string{"Party Degen"},
// 	},
// }
// rarityIndex.CalulateRarityScore(mismatchedKnightSet)

// //Rarity: 8
// halfGoldenSuit := []metadata.Attribute{
// 	{
// 		TraitType: "Eyewear",
// 		Value:     "Golden Knight Helmet",
// 		Sets:      []string{"Knight"},
// 	},
// 	{
// 		TraitType: "Footwear",
// 		Value:     "Golden Shoes",
// 		Sets:      []string{"Golden Suit"},
// 	},
// 	{
// 		TraitType: "Headwear",
// 		Value:     "Golden Hat",
// 		Sets:      []string{"Golden Suit"},
// 	},
// 	{
// 		TraitType: "Torso",
// 		Value:     "Golden Jacket",
// 		Sets:      []string{"Golden Suit"},
// 	},
// 	{
// 		TraitType: "Pants",
// 		Value:     "Golden Grieves",
// 		Sets:      []string{"Knight"},
// 	},
// 	{
// 		TraitType: "Left Hand",
// 		Value:     "Shield",
// 		Sets:      []string{"Knight"},
// 	},
// 	{
// 		TraitType: "Right Hand",
// 		Value:     "Bong",
// 		Sets:      []string{"Party Degen"},
// 	},
// }

// rarityIndex.CalulateRarityScore(halfGoldenSuit)

// //Rarity: 240
// fullGoldenSuit := []metadata.Attribute{
// 	{
// 		TraitType: "Eyewear",
// 		Value:     "Golden Sunglasses",
// 		Sets:      []string{"Golden Suit"},
// 	},
// 	{
// 		TraitType: "Footwear",
// 		Value:     "Golden Shoes",
// 		Sets:      []string{"Golden Suit"},
// 	},
// 	{
// 		TraitType: "Headwear",
// 		Value:     "Golden Hat",
// 		Sets:      []string{"Golden Suit"},
// 	},
// 	{
// 		TraitType: "Torso",
// 		Value:     "Golden Jacket",
// 		Sets:      []string{"Golden Suit"},
// 	},
// 	{
// 		TraitType: "Pants",
// 		Value:     "Golden Pants",
// 		Sets:      []string{"Golden Suit"},
// 	},
// 	{
// 		TraitType: "Left Hand",
// 		Value:     "Golden Gun",
// 		Sets:      []string{"Golden Suit"},
// 	},
// 	{
// 		TraitType: "Right Hand",
// 		Value:     "Bong",
// 		Sets:      []string{"Party Degen"},
// 	},
// }

// rarityIndex.CalulateRarityScore(fullGoldenSuit)

// //Rarity: 128
// degenSet := []metadata.Attribute{
// 	{
// 		TraitType: "Eyewear",
// 		Value:     "Bar Shades",
// 		Sets:      []string{"Party Degen"},
// 	},
// 	{
// 		TraitType: "Footwear",
// 		Value:     "Sneakers",
// 		Sets:      []string{"Party Degen"},
// 	},
// 	{
// 		TraitType: "Headwear",
// 		Value:     "Traffic Cone",
// 		Sets:      []string{"Party Degen"},
// 	},
// 	{
// 		TraitType: "Torso",
// 		Value:     "Weed Plant Tshirt",
// 		Sets:      []string{"Party Degen"},
// 	},
// 	{
// 		TraitType: "Pants",
// 		Value:     "Cargo Shorts",
// 		Sets:      []string{"Party Degen"},
// 	},
// 	{
// 		TraitType: "Left Hand",
// 		Value:     "Bong",
// 		Sets:      []string{"Party Degen"},
// 	},
// 	{
// 		TraitType: "Right Hand",
// 		Value:     "Bong",
// 		Sets:      []string{"Party Degen"},
// 	},
// }

// rarityIndex.CalulateRarityScore(degenSet)
// }

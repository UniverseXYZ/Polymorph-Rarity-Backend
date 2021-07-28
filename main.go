package main

import (
	"log"
	"os"
	"strings"

	"rarity-backend/config"
	"rarity-backend/dlt"
	"rarity-backend/handlers"
	"rarity-backend/services"
	"rarity-backend/store"
	"rarity-backend/structs"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
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

func initResources() (*dlt.EthereumClient, abi.ABI, *store.Store, string, *structs.ConfigService, structs.DBInfo) {
	// Load env variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file: " + err.Error())
	}

	// Inital step: Recover to be up to date
	ethClient := connectToEthereum()
	polymorphDBName := os.Getenv("POLYMORPH_DB")
	rarityCollectionName := os.Getenv("RARITY_COLLECTION")
	blocksCollectionName := os.Getenv("BLOCKS_COLLECTION")
	contractAddress := os.Getenv("CONTRACT_ADDRESS")
	transactionsCollectionName := os.Getenv("TRANSACTIONS_COLLECTION")
	historyCollectionName := os.Getenv("HISTORY_COLLECTION")
	morphCostCollectionName := os.Getenv("MORPH_COST_COLLECTION")

	if contractAddress == "" {
		log.Fatal("Missing contract address in .env")
	}
	if polymorphDBName == "" {
		log.Fatal("Missing polymorph db name in .env")
	}
	if rarityCollectionName == "" {
		log.Fatal("Missing rarity collection name in .env")
	}
	if blocksCollectionName == "" {
		log.Fatal("Missing block collection name in .env")
	}
	if transactionsCollectionName == "" {
		log.Fatal("Missing transactions collection name in .env")
	}
	if historyCollectionName == "" {
		log.Fatal("Missing morph history collection name in .env")
	}
	if morphCostCollectionName == "" {
		log.Fatal("Missing morph cost collection name in .env")
	}

	contractAbi, err := abi.JSON(strings.NewReader(string(store.StoreABI)))
	if err != nil {
		log.Fatal(err)
	}

	instance, err := store.NewStore(common.HexToAddress(contractAddress), ethClient.Client)
	if err != nil {
		log.Fatalln(err)
	}

	configService := config.NewConfigService("./config.json")
	dbInfo := structs.DBInfo{
		PolymorphDBName:            polymorphDBName,
		RarityCollectionName:       rarityCollectionName,
		TransactionsCollectionName: transactionsCollectionName,
		BlocksCollectionName:       blocksCollectionName,
		HistoryCollectionName:      historyCollectionName,
		MorphCostCollectionName:    morphCostCollectionName,
	}
	return ethClient, contractAbi, instance, contractAddress, configService, dbInfo
}

func main() {
	ethClient,
		contractAbi,
		instance,
		contractAddress,
		configService,
		dbInfo := initResources()

	go recoverAndPoll(
		ethClient,
		contractAbi,
		instance,
		contractAddress,
		configService,
		dbInfo)

	startAPI()
}

func startAPI() {
	// Routine two: API -> Should start after deploy?
	app := fiber.New()
	app.Get("/morphs/", handlers.GetPolymorphs)
	app.Get("/morphs/:id", handlers.GetPolymorphById)
	app.Get("/morphs/history/:id", handlers.GetPolymorphHistory)
	log.Fatal(app.Listen(8000))
}

func recoverAndPoll(ethClient *dlt.EthereumClient, contractAbi abi.ABI, store *store.Store, contractAddress string, configService *structs.ConfigService, dbInfo structs.DBInfo) {
	// Build transactions scramble transaction mapping from db
	txMap := handlers.GetTransactionsMapping(dbInfo.PolymorphDBName, dbInfo.TransactionsCollectionName)
	morphCostMap := handlers.GetMorphPriceMapping(dbInfo.PolymorphDBName, dbInfo.HistoryCollectionName)
	// Recover immediately
	services.RecoverProcess(ethClient, contractAbi, store, contractAddress, configService, dbInfo, txMap, morphCostMap)
	// Routine one: Start polling after recovery
	gocron.Every(15).Second().Do(services.RecoverProcess, ethClient, contractAbi, store, contractAddress, configService, dbInfo, txMap, morphCostMap)
	<-gocron.Start()
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
// rarityIndex.CalulateRarityScore(mismatchedSpartanSet, false)

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

// 	// ID: 9793
// 	degenSet := []metadata.Attribute{
// 		{
// 			TraitType: "Eyewear",
// 			Value:     "Bar Shades",
// 			Sets:      []string{"Party Degen"},
// 		},
// 		{
// 			TraitType: "Footwear",
// 			Value:     "Sneakers",
// 			Sets:      []string{"Party Degen"},
// 		},
// 		{
// 			TraitType: "Headwear",
// 			Value:     "Copter Hat",
// 			Sets:      []string{"Party Degen"},
// 		},
// 		{
// 			TraitType: "Torso",
// 			Value:     "Red Footbal Jersey",
// 			Sets:      []string{"Young Football Star"},
// 		},
// 		{
// 			TraitType: "Pants",
// 			Value:     "Gray Jeans",
// 			Sets:      []string{"Party Degen"},
// 		},
// 		{
// 			TraitType: "Left Hand",
// 			Value:     "Blue Degen Sword",
// 			Sets:      []string{"Party Degen"},
// 		},
// 		{
// 			TraitType: "Right Hand",
// 			Value:     "Bong",
// 			Sets:      []string{"Party Degen"},
// 		},
// 	}

// rarityIndex.CalulateRarityScore(degenSet, false)

// ID:8279
// degenSet := []metadata.Attribute{
// 	{
// 		TraitType: "Eyewear",
// 		Value:     "Monocle",
// 		Sets:      []string{"Plaid Suit", "Black Suit", "Brown Suit", "Grey Suit"},
// 	},
// 	{
// 		TraitType: "Footwear",
// 		Value:     "Silver Knight Boots",
// 		Sets:      []string{"Knight"},
// 	},
// 	{
// 		TraitType: "Headwear",
// 		Value:     "Silver Spartan Helmet",
// 		Sets:      []string{"Spartan"},
// 	},
// 	{
// 		TraitType: "Torso",
// 		Value:     "Silver Spartan Armor",
// 		Sets:      []string{"Spartan"},
// 	},
// 	{
// 		TraitType: "Pants",
// 		Value:     "Golden Grieves",
// 		Sets:      []string{"Knight"},
// 	},
// 	{
// 		TraitType: "Left Hand",
// 		Value:     "Golden Spartan Sword",
// 		Sets:      []string{"Spartan"},
// 	},
// 	{
// 		TraitType: "Right Hand",
// 		Value:     "Red Degen Sword",
// 		Sets:      []string{"Party Degen"},
// 	},
// }

// rarityIndex.CalulateRarityScore(degenSet, false)
// }

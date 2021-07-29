package db

import (
	"context"
	"log"
	"os"
	"sync"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var once sync.Once
var instance *mongo.Client

// GetDbConnection returns an instance of a mongo db client. This is a singleton pattern in order to have only one alive connection to the database.
//
// If no connection exists, it will connect to database.
//
// If connection exists, it will return the instance of the database
func GetDbConnection() *mongo.Client {
	once.Do(func() {
		client := connectToDb()
		checkConnectionAndRestore(client)
		instance = client
	})

	checkConnectionAndRestore(instance)
	return instance
}

// connectToDb retrieves db config from .env and tries to conenct to the database.
func connectToDb() *mongo.Client {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	username := os.Getenv("USERNAME")
	password := os.Getenv("PASSWORD")
	dbUrl := os.Getenv("DB_URL")

	if username == "" {
		log.Fatalln("Missing username in .env")
	}
	if password == "" {
		log.Fatalln("Missing password in .env")
	}
	if dbUrl == "" {
		log.Fatalln("Missing db url in .env")
	}

	connectionStr := "mongodb+srv://" + username + ":" + password + "@" + dbUrl + "?retryWrites=true&w=majority"

	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(connectionStr))
	if err != nil {
		log.Fatal(err)
	}

	return client
}

// checkConnectionAndRestore ping the client and it throws and error, it tries to reconnect.
func checkConnectionAndRestore(client *mongo.Client) {
	err := client.Ping(context.Background(), readpref.Primary())

	if err != nil {
		log.Fatal(err)
		newClient := connectToDb()
		client = newClient
	}
}

// GetMongoDbCollection accepts dbName and collectionname and returns an instance of the specified collection.
func GetMongoDbCollection(DbName string, CollectionName string) (*mongo.Collection, error) {
	client := GetDbConnection()

	collection := client.Database(DbName).Collection(CollectionName)
	return collection, nil
}

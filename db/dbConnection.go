package db

import (
	"context"
	"fmt"
	"log"
	"os"
	"sync"
	"time"

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
		// checkConnectionAndRestore(client)
		instance = client
	})

	// checkConnectionAndRestore(instance)
	return instance
}
const (
	// Timeout operations after N seconds
	connectTimeout  = 5
	queryTimeout    = 30
	
	// Which instances to read from
	readPreference           = "secondaryPreferred"
	connectionStringTemplate = "mongodb://%s:%s@%s/test?replicaSet=rs0&readpreference=%s&connect=direct&sslInsecure=true&retryWrites=false"
)

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

	connectionURI := fmt.Sprintf(connectionStringTemplate, username, password, dbUrl, readPreference)

	client, err := mongo.NewClient(options.Client().ApplyURI(connectionURI))
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), connectTimeout*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatalf("Failed to connect to cluster: %v", err)
	}

	// Force a connection to verify our connection string
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatalf("Failed to ping cluster: %v", err)
	}
	
	fmt.Println("Connected to DocumentDB!")

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

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

func Connect() *mongo.Client {
	once.Do(func() {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}

		username := os.Getenv("USERNAME")
		password := os.Getenv("PASSWORD")
		dbUrl := os.Getenv("DB_URL")

		connectionStr := "mongodb+srv://" + username + ":" + password + "@" + dbUrl + "?retryWrites=true&w=majority"

		client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(connectionStr))
		if err != nil {
			log.Fatal(err)
		}

		err = client.Ping(context.Background(), readpref.Primary())

		if err != nil {
			log.Fatal(err)
		}

		instance = client

	})

	return instance
}

func GetMongoDbCollection(DbName string, CollectionName string) (*mongo.Collection, error) {
	client := Connect()

	collection := client.Database(DbName).Collection(CollectionName)
	return collection, nil
}

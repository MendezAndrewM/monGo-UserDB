package connection

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// ConnectDB ...
func ConnectDB() *mongo.Collection {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Failed to load .env file")
	}

	dbName := os.Getenv("dbName")
	dbCollection := os.Getenv("dbCollectionName")

	config := GetConfiguration()
	clientOptions := options.Client().ApplyURI(config.ConnectionString)
	client, err := mongo.Connect(context.TODO(), clientOptions)

	collection := client.Database(dbName).Collection(dbCollection)

	return collection
}

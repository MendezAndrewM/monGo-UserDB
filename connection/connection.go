package connection

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/mendezandrewm/monGo-UserDB/utils"
)

// ConnectDB ...
func ConnectDB() *mongo.Client {
	utils.LoadEnvVars()

	connectionString := os.Getenv("CONNECTION_STRING")

	client, err := mongo.NewClient(options.Client().ApplyURI(connectionString))
	if err != nil {
			log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
			log.Fatal(err)
	}
	log.Println("Successfully Connected to Database")


	return client
}

var Client *mongo.Client = ConnectDB()

//OpenCollection is a  function makes a connection with a collection in the database
func OpenCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Failed to load .env file")
	}

	cluster := os.Getenv("CLUSTER_NAME")

	var collection *mongo.Collection = client.Database(cluster).Collection(collectionName)

	return collection
}

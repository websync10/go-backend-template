package db

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client
var DB *mongo.Database

func ConnectMongo() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: NO .env file found")
	}

	mongoURL := os.Getenv("MONGO_URL")
	dbName := os.Getenv("DB_NAME")

	if mongoURL == "" || dbName == "" {
		log.Fatal("❌ Missing MONGODB_URL or DB_NAME in environment variables")
	}

	clientOptions := options.Client().ApplyURI(mongoURL)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	Client, err = mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal("❌ Error connecting to MongoDB:", err)
	}

	err = Client.Ping(ctx, nil)
	if err != nil {
		log.Fatal("❌ MongoDB ping failed:", err)
	}

	DB = Client.Database(dbName)

	fmt.Println("✅ Connected to MongoDB:", dbName)
}

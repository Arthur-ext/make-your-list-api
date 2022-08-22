package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	"github.com/joho/godotenv"
)

type MongoClient struct {
	*mongo.Client
}

func NewMongoClient() (*MongoClient, error) {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("some error occured. Err: %s", err)
	}
	mongo_pass := os.Getenv("MONGO_PSWD")

	uri := fmt.Sprintf("mongodb+srv://wedding-gifts:%s@cluster0.fle2wmc.mongodb.net/?retryWrites=true&w=majority", mongo_pass)
	print(uri)
	// uri := "mongodb://root:rootpswd@mongo:27017"
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))

	dbConn := &MongoClient{
		Client: client,
	}

	return dbConn, err
}

func (m *MongoClient) CheckConn() error {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	err := m.Ping(ctx, readpref.Primary())

	return err
}

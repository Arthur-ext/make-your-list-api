package database

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type MongoClient struct {
	*mongo.Client
}

func NewMongoClient() (*MongoClient, error) {
	uri := "mongodb://root:rootpswd@localhost:27017"

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
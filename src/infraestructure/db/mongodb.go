package db

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DBHandler struct {
	MongoClient mongo.Client
	database    *mongo.Database
}

func NewDBHandler(connectStr string, dbName string) (DBHandler, error) {
	dbHandler := DBHandler{}
	clientOpts := options.Client().ApplyURI(connectStr)

	client, err := mongo.Connect(context.TODO(), clientOpts)
	if err != nil {
		log.Fatal(err)
		return dbHandler, err
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
		return dbHandler, err
	}

	dbHandler.MongoClient = *client
	dbHandler.database = client.Database(dbName)
	return dbHandler, nil
}

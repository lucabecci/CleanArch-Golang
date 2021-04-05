package db

import (
	"context"
	"log"

	"github.com/lucabecci/CleanArch-Golang/src/domain"
	"go.mongodb.org/mongo-driver/bson"
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

func (dbHandler DBHandler) FindAllAnimes() ([]*domain.Anime, error) {
	var results []*domain.Anime

	collection := dbHandler.database.Collection("animes")
	cur, err := collection.Find(context.TODO(), bson.D{{}})

	if err != nil {
		return nil, err
	}

	for cur.Next(context.TODO()) {
		var elem domain.Anime
		err2 := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err2)
		}
		results = append(results, &elem)
	}

	return results, nil
}

func (dbHandler DBHandler) SaveAnime(anime domain.Anime) error {
	collection := dbHandler.database.Collection("animes")
	_, err := collection.InsertOne(context.TODO(), anime)
	if err != nil {
		return err
	}
	return nil
}

func (dbHandler DBHandler) SaveAuthor(author domain.Author) error {
	collection := dbHandler.database.Collection("authors")
	_, err := collection.InsertOne(context.TODO(), author)
	if err != nil {
		return err
	}
	return nil
}

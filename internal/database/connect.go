package database

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var goCtx context.Context

var client *mongo.Client
var err error

func ConnectToDatabase() {

	goCtx = context.Background()

	client, err = mongo.Connect(goCtx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal("Proble In Connecting To Database")
	}

	if err := client.Ping(goCtx, nil); err != nil {
		fmt.Println("Error Connected To Database....")
	}

	articlesCollection := client.Database("articleDB").Collection("articles")
	userCollection := client.Database("articleDB").Collection("users")

	userCollection.Indexes().CreateOne(goCtx, mongo.IndexModel{
		Keys:    bson.D{{Key: "email", Value: -1}},
		Options: options.Index().SetUnique(true),
	})

	ArticleDB = *NewAticleDatabase(articlesCollection, goCtx)

	UserDB = *NewUserDatabase(userCollection, goCtx)

	fmt.Println("Connected To DB")

}

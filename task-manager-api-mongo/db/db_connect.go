package db

import (
	"context"
	"log"

	//"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Database

// function to connect to the database
func Connect()  {
	clinetOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	client, err := mongo.Connect(context.Background(), clinetOptions)

	if err != nil {
		log.Fatal(err.Error())
	}

	err = client.Ping(context.Background(), nil)
	
	if err != nil {
		log.Fatal(err.Error())
	}
	
	DB = client.Database("task_management")
	
	log.Println("Connected to MongoDB")

}


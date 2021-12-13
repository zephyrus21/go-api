package controller

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb+srv://admin-user:zephyrus21@cluster0.g7flm.mongodb.net/go-api?retryWrites=true&w=majority"

func init() {
	clientOption := options.Client().ApplyURI(connectionString)

	client, err := mongo.Connect(context.TODO(), clientOption)
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected to MongoDB!")

	collection = client.Database("go-api").Collection("go-api")

	fmt.Println("Collection established!")
}

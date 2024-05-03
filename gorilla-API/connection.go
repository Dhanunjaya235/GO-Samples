package main

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func dbConnection() {
	var err error
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	URI := GetEnvValueFromKey("MONGO_DB_URI")
	opts := options.Client().ApplyURI(URI).SetServerAPIOptions(serverAPI)

	client, err = mongo.Connect(context.TODO(), opts)

	if err != nil {
		fmt.Println("Connect Failed")
		panic(err)
	}

	database = client.Database("go-lang-poc")
	usersCollection = database.Collection("users")
	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")
}

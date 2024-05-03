package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	client          *mongo.Client
	database        *mongo.Database
	usersCollection *mongo.Collection
)

func main() {
	fmt.Println("Gorilla API POC")
	dbConnection()

	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	router := mux.NewRouter()
	userRouter := router.PathPrefix("/users").Subrouter()

	userRouter.HandleFunc("", getAllUsers).Methods("GET")
	userRouter.HandleFunc("/add", insertNewUser).Methods("POST")
	userRouter.HandleFunc("/search", findUser).Methods("GET")
	userRouter.HandleFunc("/update", updateUser).Methods("PATCH")
	userRouter.HandleFunc("/delete", deleteUser).Methods("DELETE")

	http.ListenAndServe(":9090", router)

}

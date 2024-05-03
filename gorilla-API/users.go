package main

import (
	"context"
	"encoding/json"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UsersResponse struct {
	Users  []bson.M `json:"users"`
	Status int      `json:"status"`
}

type NewUserResponse struct {
	User   bson.M `json:"user"`
	Status int    `json:"status"`
}

type UpdatedUserResponse struct {
	User   bson.M `json:"user"`
	Status int    `json:"status"`
}
type DeletedUserResponse struct {
	User   bson.M `json:"user"`
	Status int    `json:"status"`
}

func getAllUsers(w http.ResponseWriter, req *http.Request) {
	cursor, err := usersCollection.Find(context.Background(), bson.D{})

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var documents []bson.M

	for cursor.Next(context.Background()) {
		var document bson.M
		if err := cursor.Decode(&document); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		documents = append(documents, document)
	}

	var response UsersResponse

	if len(documents) == 0 {
		response = UsersResponse{Users: []bson.M{}, Status: http.StatusOK}
	} else {
		response = UsersResponse{Users: documents, Status: http.StatusOK}
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func insertNewUser(w http.ResponseWriter, req *http.Request) {

	var newUser map[string]interface{}

	if error := json.NewDecoder(req.Body).Decode(&newUser); error != nil {
		http.Error(w, error.Error(), http.StatusBadRequest)
		return
	}

	result, err := usersCollection.InsertOne(context.Background(), newUser)
	newUser["_id"] = result.InsertedID
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	response := NewUserResponse{User: newUser, Status: http.StatusOK}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func findUser(w http.ResponseWriter, req *http.Request) {

	params := req.URL.Query()

	flag := params.Get("flag")

	var filters []interface{}
	filter := make(map[string]interface{})

	for key, values := range params {
		for _, value := range values {
			paramFilter := make(map[string]interface{})
			paramFilter[key] = value
			filters = append(filters, paramFilter)
		}
	}

	if flag == "1" {
		filter["$or"] = filters
	} else {
		filter["$and"] = filters
	}

	cursor, err := usersCollection.Find(context.Background(), filter)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var documents []bson.M

	for cursor.Next(context.Background()) {
		var document bson.M
		if err := cursor.Decode(&document); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		documents = append(documents, document)
	}

	var response UsersResponse

	if len(documents) == 0 {
		response = UsersResponse{Users: []bson.M{}, Status: http.StatusOK}
	} else {
		response = UsersResponse{Users: documents, Status: http.StatusOK}
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func updateUser(w http.ResponseWriter, req *http.Request) {

	params := req.URL.Query()

	filter := make(map[string]interface{})
	_id := params.Get("_id")
	email := params.Get("email")

	if _id != "" {
		objectId, err := primitive.ObjectIDFromHex(_id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		filter["_id"] = objectId
	} else if email != "" {
		filter["email"] = email
	} else {
		http.Error(w, "You can update user only by using EMAIL or _ID", http.StatusBadRequest)
		return
	}

	updatedUser := make(map[string]interface{})
	updatedUserQuery := make(map[string]interface{})

	if error := json.NewDecoder(req.Body).Decode(&updatedUser); error != nil {
		http.Error(w, error.Error(), http.StatusBadRequest)
		return
	}
	updatedUserQuery["$set"] = updatedUser
	result, err := usersCollection.UpdateOne(context.Background(), filter, updatedUserQuery)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	updatedUser["_id"] = result.UpsertedID
	updatedUser["upsertedCount"] = result.UpsertedCount
	updatedUser["modifiedCount"] = result.ModifiedCount
	updatedUser["matchedCount"] = result.MatchedCount

	response := UpdatedUserResponse{User: updatedUser, Status: http.StatusOK}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func deleteUser(w http.ResponseWriter, req *http.Request) {

	params := req.URL.Query()

	filter := make(map[string]interface{})
	_id := params.Get("_id")
	email := params.Get("email")

	if _id != "" {
		objectId, err := primitive.ObjectIDFromHex(_id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		filter["_id"] = objectId
	} else if email != "" {
		filter["email"] = email
	} else {
		http.Error(w, "You can delete user only by using EMAIL or _ID", http.StatusBadRequest)
		return
	}

	_, err := usersCollection.DeleteOne(context.Background(), filter)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	response := DeletedUserResponse{User: filter, Status: http.StatusOK}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

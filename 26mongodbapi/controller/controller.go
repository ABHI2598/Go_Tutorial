package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	db "github.com/ABHI2598/mongoapi/database"
	"github.com/ABHI2598/mongoapi/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	collection, client, context, cancel := db.Connect()
	defer db.CloseConnection(client, context, cancel)

	movies := make([]model.Netflix, 0, 10)

	cursor, err := collection.Find(context, bson.D{{}})

	if err != nil {
		log.Fatal(err)
		http.Error(w, "Error while finding record", http.StatusNotFound)
	}

	defer cursor.Close(context)

	for cursor.Next(context) {
		var movie model.Netflix
		if err := cursor.Decode(&movie); err != nil {
			log.Fatal(err)
		}

		movies = append(movies, movie)
	}

	j, err := json.Marshal(movies)
	if err != nil {
		log.Fatal(err)
		http.Error(w, "Error while parsing result", http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

func CreateMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")
	collection, client, context, cancel := db.Connect()

	defer db.CloseConnection(client, context, cancel)

	var movie model.Netflix
	err := json.NewDecoder(r.Body).Decode(&movie)
	if err != nil {
		log.Fatal(err)
		http.Error(w, "Unable to parse request", http.StatusInternalServerError)
	}

	movie.ID = primitive.NewObjectID()
	insert, err := collection.InsertOne(context, movie)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Inserted user %v", insert.InsertedID)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(movie)

}

func UpdateOneMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "PUT,GET")
	id := r.PathValue("id")
	objectId, _ := primitive.ObjectIDFromHex(id)

	collection, client, context, cancel := db.Connect()
	defer db.CloseConnection(client, context, cancel)

	update := bson.M{"$set": bson.M{"watched": true}}

	updated, err := collection.UpdateOne(context, bson.M{"_id": objectId}, update)
	if err != nil {
		log.Fatal(err)
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(updated)

}

func GetOneMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "GET,PUT")
	id := r.PathValue("id")
	objectId, _ := primitive.ObjectIDFromHex(id)

	collection, client, context, cancel := db.Connect()
	defer db.CloseConnection(client, context, cancel)

	filter := bson.M{"_id": objectId}

	var movie model.Netflix
	err := collection.FindOne(context, filter).Decode(&movie)
	if err == mongo.ErrNoDocuments {
		http.Error(w, "No Document found with given id", http.StatusNotFound)
	} else if err != nil {
		log.Fatal(err)
	}

	w.WriteHeader(http.StatusFound)
	json.NewEncoder(w).Encode(movie)

}

func DeleteOneMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")
	id := r.PathValue("id")

	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Fatal(err)
	}

	collection, client, context, cancel := db.Connect()
	defer db.CloseConnection(client, context, cancel)

	filter := bson.M{"_id": objectId}

	result, err := collection.DeleteOne(context, filter)
	if err != nil {
		log.Fatal(err)
	}

	resultStr := fmt.Sprintf("Movie deleted %v with id of the movie %v", result.DeletedCount, objectId)

	w.WriteHeader(http.StatusNoContent)
	fmt.Fprintln(w, resultStr)

}

package mongodb

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var connectionString = "mongodb+srv://<user>:<password>@cluster0.7fs7f.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0"
var dbName = "dbname"
var collectionName = "dbcollection"

func Connect() (*mongo.Collection, *mongo.Client, context.Context, context.CancelFunc) {
	ctx, cancel := context.WithTimeout(context.TODO(), 30*time.Second)

	option := options.Client().ApplyURI(connectionString)
	client, err := mongo.Connect(ctx, option)

	if err != nil {
		log.Fatal(err)
	}

	collection := client.Database(dbName).Collection(collectionName)
	return collection, client, ctx, cancel
}

func CloseConnection(client *mongo.Client, context context.Context, cancel context.CancelFunc) {
	defer func() {
		cancel()
		if err := client.Disconnect(context); err != nil {
			log.Fatal(err)
		}
		fmt.Println("Connection is closed")
	}()
}

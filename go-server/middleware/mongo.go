package middleware

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var dbName string = "myFirstDatabase"

type Auth struct {
	Username string
	Password string
}

// getAuth get mongodb auth
func getAuth() Auth {
	data, err := os.Open("./crawling/secret/mongodb_auth.json")
	check(err)
	var auth Auth
	byteValue, _ := ioutil.ReadAll(data)
	json.Unmarshal(byteValue, &auth)
	return auth
}

// connectionDBInstance connect Mongo DB
func ConnectionDBInstance(colName string) *mongo.Collection {
	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		log.Fatal("You must set your 'MONGODB_URI' environmental variable.")
	}
	Authorization := getAuth()

	client, err := mongo.Connect(context.TODO(),
		options.Client().ApplyURI(uri).SetAuth(options.Credential{
			Username: Authorization.Username,
			Password: Authorization.Password,
		}))
	check(err)

	// Ping the primary
	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}

	collection := client.Database(dbName).Collection(colName)
	fmt.Println("Collection instance connect")

	return collection
}

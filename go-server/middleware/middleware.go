package middleware

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var dbName string = "myFirstDatabase"
var colName string = "hdfs"

var collection *mongo.Collection

// create connection to mongo db
func init() {
	connectionDBInstance()
}

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
func connectionDBInstance() {
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

	collection = client.Database(dbName).Collection(colName)
	fmt.Println("Collection instance connect")

}

func GetAllConfig(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	results := getAllTask()
	fmt.Println(results)
}

func getAllTask() []primitive.M {
	cur, err := collection.Find(context.Background(), bson.D{{}})
	check(err)

	var results []primitive.M
	for cur.Next(context.Background()) {
		var result bson.M
		err = cur.Decode(&result)
		check(err)
		results = append(results, result)
	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
	cur.Close(context.Background())
	return results
}

func check(e error) {
	if e != nil {
		fmt.Println(e)
		panic(e)
	}
}

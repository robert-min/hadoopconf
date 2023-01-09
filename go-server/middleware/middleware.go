package middleware

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// GetAllHdfsConfig get Hdfs config data API
func GetAllHdfsConfig(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	configs := getHdfsConfigDB()
	json.NewEncoder(w).Encode(configs)
}

// getHdfsConfigDB get Hdfs config data on mongodb
func getHdfsConfigDB() []primitive.M {
	collection := ConnectionDBInstance("hdfs")
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

// GetAllHdfsConfig get Hdfs config data API
func GetAllCoreConfig(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	configs := getCoreConfigDB()
	json.NewEncoder(w).Encode(configs)
}

// getHdfsConfigDB get Hdfs config data on mongodb
func getCoreConfigDB() []primitive.M {
	collection := ConnectionDBInstance("core")
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

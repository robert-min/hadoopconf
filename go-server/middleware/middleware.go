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
	configs := getConfigDB("hdfs")
	json.NewEncoder(w).Encode(configs)
}

// GetAllCoreConfig get Core config data API
func GetAllCoreConfig(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	configs := getConfigDB("core")
	json.NewEncoder(w).Encode(configs)
}

// getConfigDB get colName config data on mongodb
func getConfigDB(colName string) []primitive.M {
	collection := ConnectionDBInstance(colName)
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

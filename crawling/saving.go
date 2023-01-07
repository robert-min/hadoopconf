package crawling

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"

	strip "github.com/grokify/html-strip-tags-go"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var dbName string = "config"
var colName string = "hdfs"

// connectionMongo to connect DB
func connectionMongo() (client *mongo.Client, ctx context.Context, cancel context.CancelFunc) {

	ctx, cancel = context.WithTimeout(context.Background(), 3*time.Second)

	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		log.Fatal("You must set your 'MONGODB_URI' environmental variable.")
	}

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	check(err)
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
	// Ping the primary
	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}

	return client, ctx, cancel
}

// GetCollection to connect collection
func GetCollection(client *mongo.Client, colName string) *mongo.Collection {
	return client.Database(dbName).Collection(colName)
}

func createConfig(name string, value string, description string) {
	// client, ctx, cancel := connectionMongo()
	// defer client.Disconnect(ctx)
	// defer cancel()

	// num, err := GetCollection(client, colName).CountDocuments(ctx, bson.M{"name": name})
	// check(err)

	// GetCollection(client, colName).InsertOne()

	return nil
}

type config struct {
	Name        string
	Value       string
	Description string
}

func Parsing(path string) {
	data, err := ioutil.ReadFile(path)
	check(err)
	docs := strings.Split(string(data), "\n\n")

	for _, context := range docs {
		lines := strings.Split(context, "\n")
		var b bytes.Buffer
		configMap := new(config)

		for _, line := range lines {
			line = strings.TrimLeft(line, " ")
			if strings.HasPrefix(line, "<name>") {
				name := strip.StripTags(line)
				configMap.Name = name
			} else if strings.HasPrefix(line, "<value>") {
				value := strip.StripTags(line)
				configMap.Value = value
			} else if strings.HasPrefix(line, "</property>") {
				description := b.String()
				description = strip.StripTags(description)
				configMap.Description = description
				createConfig(configMap.Name, configMap.Value, configMap.Description)

				configMap = new(config)
				b.Reset()
			} else {
				b.WriteString(line)
			}
		}

	}

}
func check(e error) {
	if e != nil {
		fmt.Println(e)
		panic(e)
	}
}

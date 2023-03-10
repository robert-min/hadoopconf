package crawling

import (
	"bytes"
	"context"
	"encoding/json"
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

var dbName string = "myFirstDatabase"
var colName string = "core"

type config struct {
	Name        string
	Value       string `bson:"value,omitempty"`
	Description string `bson:"description,omitempty"`
}

func ParsingCreate(path string) {
	data, err := ioutil.ReadFile(path)
	check(err)
	docs := strings.Split(string(data), "\n\n")
	allDatas := []interface{}{}
	var name string = ""
	var value string = ""
	var description string = ""
	var b bytes.Buffer
	for _, context := range docs {
		lines := strings.Split(context, "\n")

		for _, line := range lines {
			line = strings.TrimLeft(line, " ")
			if strings.HasPrefix(line, "<name>") {
				name = strip.StripTags(line)
			} else if strings.HasPrefix(line, "<value>") {
				value = strip.StripTags(line)
			} else if strings.HasPrefix(line, "</property>") {
				description = b.String()
				description = strip.StripTags(description)
				allDatas = append(allDatas, config{Name: name, Value: value, Description: description})
				b.Reset()
			} else {
				b.WriteString(line)
			}
		}
	}
	client, ctx, cancel := connectionMongo()
	defer client.Disconnect(ctx)
	defer cancel()

	coll := client.Database(dbName).Collection(colName)
	result, err := coll.InsertMany(context.TODO(), allDatas)
	check(err)
	fmt.Printf("%d documents inserted with IDs:\n", len(result.InsertedIDs))
	for _, id := range result.InsertedIDs {
		fmt.Printf("\t%s\n", id)
	}
}

type Auth struct {
	Username string
	Password string
}

func getAuth() Auth {
	data, err := os.Open("./crawling/secret/mongodb_auth.json")
	check(err)
	var auth Auth
	byteValue, _ := ioutil.ReadAll(data)
	json.Unmarshal(byteValue, &auth)
	return auth
}

// connectionMongo to connect DB
func connectionMongo() (client *mongo.Client, ctx context.Context, cancel context.CancelFunc) {

	ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

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
	// defer func() {
	// 	if err := client.Disconnect(context.TODO()); err != nil {
	// 		panic(err)
	// 	}
	// }()
	// Ping the primary
	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}

	return client, ctx, cancel
}

func check(e error) {
	if e != nil {
		fmt.Println(e)
		panic(e)
	}
}

package crawling

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// Connection URI

func connectionMongo() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		log.Fatal("You must set your 'MONGODB_URI' environmental variable. See\n\t https://www.mongodb.com/docs/drivers/go/current/usage-examples/#environment-variable")
	}
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

}

func Parsing(path string) {
	connectionMongo()
	data, err := ioutil.ReadFile(path)
	check(err)
	docs := strings.Split(string(data), "\n\n")
	for _, context := range docs {
		fmt.Println("=============")
		lines := strings.Split(context, "\n")
		// for _, line := range lines {
		// 	if strings.HasPrefix(line, "<name>") {

		// 	}
		// }

	}

}

package main

import (
	"context"
	"math/rand"
	"net/http"
	"os"
	"time"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.mongodb.org/mongo-driver/v2/mongo/readpref"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func randomString(length int) string {
	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))
	result := make([]byte, length)
	for i := range result {
		result[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(result)
}

type Link struct {
	ID string `bson:"id"`
	RedirectTo string `bson:"redirect_to"`
	CreatedAt time.Time `bson:"created_at"`
}

func newLink(redirectTo string) *Link {
	return &Link{
		ID: randomString(10),
		RedirectTo: redirectTo,
		CreatedAt: time.Now(),
	}
}

func fetch(url string, valid chan bool) {
	resp, err := http.Get(url)

	if err != nil {
		valid <- false
		return
	}

	defer resp.Body.Close()
	valid <- true
}

func main() {
	mongodb_uri := os.Getenv("MONGODB_URI")
	db_name := os.Getenv("DB_NAME")
	collection_name := os.Getenv("COLLECTION_NAME")
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
    clientOpts := options.Client().ApplyURI(mongodb_uri).SetServerAPIOptions(serverAPI)

	client, err := mongo.Connect(clientOpts)

	if (err != nil) {
		panic(err)
	}

	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}

	collection := client.Database(db_name).Collection(collection_name)

	router := gin.Default()
	router.POST("/createLink", func(ctx *gin.Context) {
		redirectTo := ctx.Query("redirectTo")

		if (redirectTo == "") {
			ctx.JSON(400, gin.H{
				"error": "no redirect link found",
			})
			return
		}

		fetchChannel := make(chan bool)
		go fetch(redirectTo, fetchChannel)
		valid := <-fetchChannel
		link := newLink(redirectTo)
		if (valid) {
			_, err := collection.InsertOne(context.TODO(), link)
			if err != nil {
				ctx.JSON(500, gin.H{
					"error": err.Error(),
				})
				return
			}
			ctx.JSON(201, gin.H{
				"id": link.ID,
			})
			return
		} else {
			ctx.JSON(400, gin.H{
				"error": "invalid link",
			})
			return
		}
	})

	router.Run()

}
package db


import (
    "GoDiscordBot/config"
    "context"
    "fmt"
    "log"
    "time"

    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"

type User struct {
	Id      string
	Name    string
	Balance int
	}
	
	
	var client *mongo.Client
	var GoDB *mongo.Database
	var UsersCollection *mongo.Collection
	
	
	func Connect() {
	 	serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)
		clientOptions := options.Client().ApplyURI(config.MongoDBURL).SetServerAPIOptions(serverAPIOptions)
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		client, err := mongo.Connect(ctx, clientOptions)
		if err != nil {
			log.Fatal(err)
		}
	
	
		GoDB = client.Database("go")
		UsersCollection = GoDB.Collection("users")
	
	
		err = client.Ping(ctx, readpref.Primary())
		if err != nil {
			log.Fatal(err)
		}
		databases, err := client.ListDatabaseNames(ctx, bson.M{})
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(databases)
	}
	
	func Disconnect() {
		ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
		client.Disconnect(ctx)
	}
	

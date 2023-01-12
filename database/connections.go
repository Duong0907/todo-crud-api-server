package db

import (
	"context"
	"fmt"
	"time"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var Client *mongo.Client
var ctx context.Context
var cancel context.CancelFunc

func Connect() (error) {
	
	if err := godotenv.Load(".env"); err != nil {
		return err
	}
	mongodbURL := os.Getenv("MONGODB_URL")
	fmt.Println(mongodbURL)
	ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
	var err error
	Client, err = mongo.NewClient(options.Client().ApplyURI(mongodbURL))
	if err != nil {
		return err
	}

	err = Client.Connect(ctx)
	return err
}

func Close() {
	defer cancel()
	defer func() {
		if err := Client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()
}

func Ping() error {
	if err := Client.Ping(ctx, readpref.Primary()); err != nil {
		return err
	}
	fmt.Println("Connect successfully!")
	return nil
}

func openCollection(collectionName string) *mongo.Collection {
	return Client.Database("cluster0").Collection(collectionName)
}

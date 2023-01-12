package db

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var Client *mongo.Client
var ctx context.Context

var cancel context.CancelFunc

func Connect() (error) {
	const uri = "mongodb+srv://Duong:09072003Tpl@cluster0.aszqp6w.mongodb.net/?retryWrites=true&w=majority"

	ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
	var err error
	Client, err = mongo.NewClient(options.Client().ApplyURI(uri))
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

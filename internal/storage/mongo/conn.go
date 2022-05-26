package mongo

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewConn(host, port, db string) (*mongo.Database, context.CancelFunc) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	clientOpts := options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%s", host, port))
	client, err := mongo.Connect(ctx, clientOpts)
	if err != nil {
		cancel()
		log.Fatal(err)
	}
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	database := client.Database(db)
	return database, cancel
}

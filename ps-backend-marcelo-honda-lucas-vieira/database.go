package main

import (
	"context"
	"fmt"
	"time"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)	


func ConnectMongoDB() (*mongo.Collection, context.Context, context.CancelFunc, *mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
	
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://127.0.0.1/27017"))
	if err != nil {
		cancel()
		log.Printf("Failed to connect to MongoDB: %v", err)
		return nil, nil, nil, nil, fmt.Errorf("Failed to connect to MongoDB: %v", err)
	}
	
	collection := client.Database("produtos-mhk-lvp").Collection("products")
	return collection, ctx, cancel, client, nil
}

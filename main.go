package main

import (
	"context"
	"fmt"
	"log"
	"oms/database"
	"oms/listener"
	"oms/routes"
	"time"

	appinit "oms/init"

	"github.com/omniful/go_commons/http"
	"go.mongodb.org/mongo-driver/bson"
)

func main() {
	ctx := context.Background()

	appinit.Initialize(ctx)

	runHttpServer(ctx)

	// Check if the MongoDB connection is valid
	if database.DB == nil {
		log.Fatal("MongoDB client is not initialized. Exiting...")
	}

	// Example check for DB connection
	collection := database.DB.Database("OMS").Collection("orders")
	count, err := collection.CountDocuments(ctx, bson.D{})
	if err != nil {
		log.Fatal("Error fetching orders count:", err)
	}
	fmt.Println("Total Orders in DB:", count)

	// Start SQS consumer in a separate goroutine
	go listener.StartConsume()

}

func runHttpServer(ctx context.Context) {
	server := http.InitializeServer(":8081", 10*time.Second, 10*time.Second, 70*time.Second)

	if err := server.StartServer("OMS Service Started"); err != nil {
		log.Fatal("Error starting the server:", err)
	}

	err := routes.PublicRoutes(ctx, server)
	if err != nil {
		log.Fatal("Error setting up routes:", err)
		return
	}

}

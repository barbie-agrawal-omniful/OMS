package database

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/omniful/go_commons/sqs"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Client

var Queue *sqs.Queue

func getDatabaseUri() string {
	return "mongodb://127.0.0.1:27017/OMS"
}

func ConnectMongo(c context.Context) {
	fmt.Println("Connecting to mongo...")
	ctx, cancel := context.WithTimeout(c, 10*time.Second)
	defer cancel()

	clientOptions := options.Client().ApplyURI(getDatabaseUri())

	var err error

	DB, err = mongo.Connect(ctx, clientOptions)

	if err != nil {
		fmt.Println("Error connecting to MongoDB:", err)
		return
	}

	err = DB.Ping(ctx, nil)
	if err != nil {
		fmt.Println("Failed to ping MongoDB:", err)
		return
	}

	fmt.Println("Successfully connected to MongoDB!")
}

func ConnectSqs(ctx context.Context) {
	if err := godotenv.Load(); err != nil {
		fmt.Println("Error loading .env file:", err)
	}
	acc := os.Getenv("AWS_ACCOUNT")
	queue_name := os.Getenv("QUEUE_NAME")
	fmt.Print("acc -> ", acc, "\n")
	fmt.Print("queue name -> ", queue_name, "\n")

	sqsConfig := sqs.GetSQSConfig(ctx, false, "ord", "eu-north-1", acc, "")
	fmt.Println(acc)
	fmt.Println("acc")
	url, err := sqs.GetUrl(ctx, sqsConfig, queue_name)
	fmt.Println(*url)
	if err != nil {
		fmt.Println(err)
	}
	queueInstance, err := sqs.NewStandardQueue(ctx, queue_name, sqsConfig)
	Queue = queueInstance
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(queueInstance)

	// this will send a message to queue in sqs
	// orders.SetProducer(ctx, queueInstance)

}

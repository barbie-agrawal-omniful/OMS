package appinit

import (
	"context"
	"oms/database"
	oms_kafka "oms/kafkaaaa"
)

func Initialize(ctx context.Context) {
	initializeDb(ctx)
	initializeSqs(ctx)
	go oms_kafka.InitializeKafkaConsumer(ctx)
	oms_kafka.InitializeKafkaProducer()
}

func initializeDb(ctx context.Context) {
	database.ConnectMongo(ctx)
}

func initializeSqs(ctx context.Context) {
	database.ConnectSqs(ctx)
}

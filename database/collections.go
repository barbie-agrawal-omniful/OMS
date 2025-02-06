package database

import (
	"go.mongodb.org/mongo-driver/mongo"
)

func GetOrdersCollection() *mongo.Collection {
	return DB.Database("OMS").Collection("orders")
}

func GetOrderItemsCollection() *mongo.Collection {
	return DB.Database("OMS").Collection("order_items")
}

func GetOrderRequestsCollection() *mongo.Collection {
	return DB.Database("OMS").Collection("order_requests")
}

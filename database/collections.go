package database

import (
	"context"
	"fmt"
	"oms/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
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

// GetFilteredOrders fetches orders based on filters
func GetFilteredOrders(filters map[string]string) ([]models.Orders, error) {
	// A query filter
	query := bson.M{}

	// Apply filters if they exist
	if tenantID, ok := filters["tenant_id"]; ok {
		query["tenant_id"] = tenantID
	}
	if sellerID, ok := filters["seller_id"]; ok {
		query["seller_id"] = sellerID
	}
	if status, ok := filters["status"]; ok {
		query["status"] = status
	}
	if startDateStr, ok := filters["start_date"]; ok {
		if endDateStr, ok := filters["end_date"]; ok {
			startDate, err := time.Parse(time.RFC3339, startDateStr)
			if err != nil {
				return nil, fmt.Errorf("invalid start_date format: %v", err)
			}
			endDate, err := time.Parse(time.RFC3339, endDateStr)
			if err != nil {
				return nil, fmt.Errorf("invalid end_date format: %v", err)
			}
			// Filter orders created within the date range
			query["created_at"] = bson.M{
				"$gte": startDate,
				"$lte": endDate,
			}
		}
	}

	// Get the orders collection
	collection := GetOrdersCollection()

	// Execute the query
	cursor, err := collection.Find(context.TODO(), query)
	if err != nil {
		return nil, fmt.Errorf("error querying orders: %v", err)
	}
	defer cursor.Close(context.TODO())

	// Decode results into a slice of orders
	var orders []models.Orders
	if err := cursor.All(context.TODO(), &orders); err != nil {
		return nil, fmt.Errorf("error decoding orders: %v", err)
	}

	return orders, nil
}

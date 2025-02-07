package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Orders struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	OrderNumber  string             `bson:"order_number" json:"order_number"`
	CustomerName string             `bson:"customer_name,omitempty" json:"customer_name"`
	Status       string             `bson:"status" json:"status"`
	TotalAmount  float64            `bson:"total_amount" json:"total_amount"`
	CreatedAt    time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt    time.Time          `bson:"updated_at" json:"updated_at"`
	TenantID     string             `bson:"tenant_id" json:"tenant_id"`
	SellerID     string             `bson:"seller_id" json:"seller_id"`
	OrderItem    []OrderItems       `bson:"order_items" json:"order_items"`
	OrderNo      string             `bson:"order_no" json:"order_no"`
}

type OrderItems struct {
	SKUID    string `bson:"sku_id,omitempty" json:"sku_id"`
	Quantity int    `bson:"quantity" json:"quantity"`
}

type KafkaResponseOrderMessage struct {
	OrderItemsID    string `json:"order_items_id"`
	OrderID         string `json:"OrderID"`
	SKUID           string `json:"sku_id"`
	QuantityOrdered int    `json:"quantity_ordered"`
	HubID           string `json:"hub_id"`
	SellerID        string `json:"seller_id"`
}

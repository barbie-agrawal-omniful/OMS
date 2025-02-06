package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Orders struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	OrderNumber string             `bson:"order_number" json:"order_number"`
	CustomerID  string             `bson:"customer_id" json:"customer_id"`
	Status      string             `bson:"status" json:"status"`
	TotalAmount float64            `bson:"total_amount" json:"total_amount"`
	CreatedAt   time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt   time.Time          `bson:"updated_at" json:"updated_at"`
}

type OrderItems struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	OrderID     primitive.ObjectID `bson:"order_id" json:"order_id"`
	ProductID   string             `bson:"product_id" json:"product_id"`
	ProductName string             `bson:"product_name" json:"product_name"`
	Quantity    int                `bson:"quantity" json:"quantity"`
	Price       float64            `bson:"price" json:"price"`
	CreatedAt   time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt   time.Time          `bson:"updated_at" json:"updated_at"`
}

type OrderRequest struct {
	HubId    string `bson:"hub_id" json:"hub_id"`
	TenantID string `bson:"tenant_id" json:"tenant_id"`
	Path     string `bson:"csv_path" json:"csv_path"`
}

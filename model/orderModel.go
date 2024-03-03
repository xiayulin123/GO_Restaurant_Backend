package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Order struct {
	ID         primitive.ObjectID `bson:"_id"`
	Order_Date time.Time          `json:"order_date" validate:"required"`
	Order_id   string             `json:"order_id"`
	Table_id   *string            `json:"table_id" validate:"required"`
	Created_at time.Time          `json:"created_at"`
	Updated_at time.Time          `json:"updated_at"`
}

package model

import (
	"time"
)

type Gift struct {
	ID        interface{} `json:"id" bson:"_id"`
	Name      string      `json:"name"`
	Assigned  string      `json:"assigned"`
	CreatedAt time.Time   `json:"createdAt" bson:"createdAt"`
	UpdatedAt time.Time   `json:"updatedAt" bson:"updatedAt"`
}

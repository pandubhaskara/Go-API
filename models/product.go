package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	// CollectionArticle holds the name of the articles collection
	CollectionProduct = "products"
)

type Product struct {
	Id       primitive.ObjectID `json:"id" bson:"_id"`
	Name     string             `json:"name" bson:"name"`
	Stock    int                `json:"stock" bson:"stock"`
	Price    int                `json:"price" bson:"price"`
	Category string             `json:"category" bson:"category"`
}

package mongo

import "go.mongodb.org/mongo-driver/bson/primitive"

type IEntity struct {
	Id primitive.ObjectID `json:"id" bson:"_id"`
}

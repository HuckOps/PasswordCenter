package tag

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Taget struct {
	ID primitive.ObjectID `bson:"_id"`
	Tag string `bson:"tag"`
}

type AddTagStruct struct {
	Tag string `json:"tag"`
}

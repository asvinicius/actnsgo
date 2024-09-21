package super

import "go.mongodb.org/mongo-driver/bson/primitive"

type Super struct {
	ID       primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name     string             `json:"name" bson:"name"`
	Login    string             `json:"login" bson:"login"`
	Password string             `json:"password" bson:"password"`
	Status   bool               `json:"status" bson:"status"`
}

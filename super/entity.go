package super

import "go.mongodb.org/mongo-driver/bson/primitive"

type Super struct {
	Super_id       primitive.ObjectID `json:"super_id,omitempty" bson:"super_id,omitempty"`
	Super_name     string             `json:"super_name" bson:"super_name"`
	Super_login    string             `json:"super_login" bson:"super_login"`
	Super_password string             `json:"super_password" bson:"super_password"`
	Super_status   bool               `json:"super_status" bson:"super_status"`
}

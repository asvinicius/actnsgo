package db

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Insert(collection string, data any) (primitive.ObjectID, error) {
	client, ctx := getConnection()
	defer client.Disconnect(ctx)

	c := client.Database(dbname).Collection(collection)

	response, err := c.InsertOne(context.Background(), data)

	if err != nil {
		return primitive.NewObjectID(), err
	}

	return response.InsertedID.(primitive.ObjectID), nil
}

package booklog

import (
	"context"
	"grpc-server/src/infra/mongodb/mcoll"
	"reflect"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var reflectType = reflect.TypeOf(&BookDeleteHistory{})

const (
	// ThisCollName .
	ThisCollName = "DELETE_LOG"
)

// BookDeleteHistory .
type BookDeleteHistory struct {
	ID   primitive.ObjectID `bson:"_id,omitempty"`
	Name string             `bson:"name"`
	User string             `bson:"user"`
}

func getColl(ctx context.Context) *mcoll.MColl {
	return mcoll.New(ctx, "TEST_DB", ThisCollName)
}

// Insert .
func Insert(ctx context.Context, name string, user string) interface{} {

	bookDeleteHistory := &BookDeleteHistory{
		Name: name,
		User: user,
	}

	return getColl(ctx).InsertOne(bookDeleteHistory)
}

// Get .
func Get(ctx context.Context, name string) *BookDeleteHistory {

	query := bson.M{"name": name}
	result := getColl(ctx).FindOne(query, reflectType)
	if result == nil {
		return nil
	}

	return result.(*BookDeleteHistory)
}

// GetAll .
func GetAll(ctx context.Context) []*BookDeleteHistory {
	query := bson.M{}
	result := getColl(ctx).Find(query, reflectType)
	if result == nil {
		return nil
	}
	return result.([]*BookDeleteHistory)
}

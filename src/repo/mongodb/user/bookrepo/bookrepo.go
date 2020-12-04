package bookrepo

import (
	"context"
	"grpc-server/src/infra/mongodb/mcoll"
	"reflect"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var reflectType = reflect.TypeOf(&UserBook{})

const (
	// ThisCollName .
	ThisCollName = "USER_BOOK"
)

// UserBook .
type UserBook struct {
	ID    primitive.ObjectID `bson:"_id,omitempty"`
	Name  string             `bson:"name"`
	Price int32              `bson:"price"`
	User  string             `bson:"user"`
}

func getColl(ctx context.Context) *mcoll.MColl {
	return mcoll.New(ctx, "TEST_DB", ThisCollName)
}

// Insert .
func Insert(ctx context.Context, name string, price int32, user string) interface{} {

	userBook := &UserBook{
		Name:  name,
		Price: price,
		User:  user,
	}

	return getColl(ctx).InsertOne(userBook)
}

// UpdatePrice .
func UpdatePrice(ctx context.Context, name string, price int32) int32 {
	query := bson.M{"name": name}

	update := bson.M{
		"$set": bson.M{
			"price": price,
		},
	}

	return getColl(ctx).UpdateOne(query, update)
}

// Delete .
func Delete(ctx context.Context, name string) int32 {
	query := bson.M{"name": name}

	return getColl(ctx).DeleteOne(query)
}

// Get .
func Get(ctx context.Context, name string) *UserBook {

	query := bson.M{"name": name}
	result := getColl(ctx).FindOne(query, reflectType)
	if result == nil {
		return nil
	}

	return result.(*UserBook)
}

// GetAll .
func GetAll(ctx context.Context) []*UserBook {
	query := bson.M{}
	result := getColl(ctx).Find(query, reflectType)
	if result == nil {
		return nil
	}
	return result.([]*UserBook)
}

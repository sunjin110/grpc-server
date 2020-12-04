package banklog

import (
	"context"
	"grpc-server/src/infra/mongodb/mcoll"
	"reflect"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var reflectType = reflect.TypeOf(&UserCashFlow{})

const (
	// ThisCollName .
	ThisCollName = "BANK_ACTION"
)

// UserCashFlow .
type UserCashFlow struct {
	ID     primitive.ObjectID `bson:"_id,omitempty"`
	User   string             `bson:"user"`
	Action string             `bson:"action"`
	Money  int32              `bson:"money"`
}

func getColl(ctx context.Context) *mcoll.MColl {
	return mcoll.New(ctx, "TEST_DB", ThisCollName)
}

// Insert .
func Insert(ctx context.Context, user string, action string, money int32) interface{} {

	userCashFlow := &UserCashFlow{
		User:   user,
		Action: action,
		Money:  money,
	}

	return getColl(ctx).InsertOne(userCashFlow)
}

// GetAll .
func GetAll(ctx context.Context) []*UserCashFlow {
	query := bson.M{}
	result := getColl(ctx).Find(query, reflectType)
	if result == nil {
		return nil
	}
	return result.([]*UserCashFlow)
}

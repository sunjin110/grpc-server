package bankrepo

import (
	"context"
	"grpc-server/src/infra/mongodb/mcoll"
	"reflect"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var reflectType = reflect.TypeOf(&BankAccount{})

const (
	// ThisCollName .
	ThisCollName = "BANK_ACCOOUNT"
)

// BankAccount .
type BankAccount struct {
	ID    primitive.ObjectID `bson:"_id,omitempty"`
	User  string             `bson:"user"`
	Money int32              `bson:"money"`
}

func getColl(ctx context.Context) *mcoll.MColl {
	return mcoll.New(ctx, "TEST_DB", ThisCollName)
}

// Insert 口座の追加
func Insert(ctx context.Context, user string) interface{} {

	account := &BankAccount{
		User:  user,
		Money: 0,
	}

	return getColl(ctx).InsertOne(account)
}

// Delete 口座の削除
func Delete(ctx context.Context, user string) int32 {
	query := bson.M{"user": user}

	return getColl(ctx).DeleteOne(query)
}

// IncUserMoney 入金
func IncUserMoney(ctx context.Context, user string, money int32) int32 {
	query := bson.M{"user": user}

	update := bson.M{
		"$inc": bson.M{
			"money": money,
		},
	}

	return getColl(ctx).UpdateOne(query, update)
}

// DecUserMoney 出金
func DecUserMoney(ctx context.Context, user string, money int32) int32 {
	query := bson.M{"user": user}

	update := bson.M{
		"$inc": bson.M{
			"money": -money,
		},
	}

	return getColl(ctx).UpdateOne(query, update)
}

// TransferUserMoney 送金
func TransferUserMoney(ctx context.Context, userfrom string, userto string, money int32) int32 {
	query := bson.M{"userfrom": userfrom}

	update := bson.M{
		"$inc": bson.M{
			"money": money,
		},
	}

	// userfrom : $inc -money

	// userto : $inc money

	return getColl(ctx).UpdateOne(query, update)
}

// Get .
func Get(ctx context.Context, user string) *BankAccount {

	query := bson.M{"user": user}
	result := getColl(ctx).FindOne(query, reflectType)
	if result == nil {
		return nil
	}

	return result.(*BankAccount)
}

// GetAll .
func GetAll(ctx context.Context) []*BankAccount {
	query := bson.M{}
	result := getColl(ctx).Find(query, reflectType)
	if result == nil {
		return nil
	}
	return result.([]*BankAccount)
}

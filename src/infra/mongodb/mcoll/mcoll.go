package mcoll

import (
	"context"
	"grpc-server/src/comm/refutil"
	"grpc-server/src/infra/mongodb/mconn"
	"reflect"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MColl .
type MColl struct {
	col *mongo.Collection
	ctx context.Context
}

// NewDefault .
func NewDefault(ctx context.Context) *MColl {
	return New(ctx, "TEST_DB", "TEST_COLL")
}

// New GetCollection
func New(ctx context.Context, dbName string, colName string) *MColl {
	return &MColl{
		col: mconn.Client.Database(dbName).Collection(colName),
		ctx: ctx,
	}
}

// InsertOne 戻り値は_id
func (c *MColl) InsertOne(data interface{}) interface{} {
	res, err := c.col.InsertOne(c.ctx, data)
	if err != nil {
		panic(err)
	}
	return res.InsertedID
}

// UpdateOne .
func (c *MColl) UpdateOne(query bson.M, update bson.M) int32 {
	res, err := c.col.UpdateOne(c.ctx, query, update)
	if err != nil {
		panic(err)
	}
	return int32(res.ModifiedCount)
}

// DeleteOne .
func (c *MColl) DeleteOne(query bson.M) int32 {
	res, err := c.col.DeleteOne(c.ctx, query)
	if err != nil {
		panic(err)
	}
	return int32(res.DeletedCount)
}

// FindOne .
func (c *MColl) FindOne(query bson.M, reflectType reflect.Type) interface{} {
	res := c.col.FindOne(c.ctx, query)
	return checkAndReflectSingleResult(res, reflectType)
}

// Find .
func (c *MColl) Find(query bson.M, reflectType reflect.Type) interface{} {
	reflectSlice := refutil.CreateSlice(reflectType)
	return reflectionFind(c.ctx, c.col, query, reflectSlice, reflectType)
}

func checkAndReflectSingleResult(result *mongo.SingleResult, reflectType reflect.Type) interface{} {
	// 検索がヒットしない場合、エラーとなる。その場合、nilを返す。
	if result.Err() != nil && result.Err().Error() == "mongo: no documents in result" {
		return nil // TODO ↑のエラーメッセージが変更される危険性があるため、テストを書いておいたほうが良いかも。
	}

	if result.Err() != nil {
		panic(result.Err())
	}

	v := reflect.New(reflectType)
	err := result.Decode(v.Interface())
	if err != nil {
		panic(err)
	}

	return v.Elem().Interface()
}

// リフレクションタイプでデータを取得する
func reflectionFind(ctx context.Context, collection *mongo.Collection, query interface{}, reflectSlice reflect.Value, reflectType reflect.Type, opts ...*options.FindOptions) interface{} {
	cur, err := collection.Find(ctx, query, opts...)
	defer func() {
		if cur != nil {
			cur.Close(ctx)
		}
		if err2 := recover(); err2 != nil {
			panic(err2)
		}
	}()
	if err != nil {
		panic(err)
	}

	for cur.Next(ctx) {
		v := reflect.New(reflectType)
		err = cur.Decode(v.Interface())
		if err != nil {
			panic(err)
		}
		reflectSlice.Set(reflect.Append(reflectSlice, v.Elem()))
	}

	err = cur.Err()
	if err != nil {
		panic(err)
	}

	return reflectSlice.Interface()
}

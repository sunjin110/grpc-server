// Package mconn MongoCollection
package mconn

import (
	"context"
	"grpc-server/src/comm/refutil"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// Client MongoDB Client
var Client *mongo.Client

// Setup MongoDB Connect
func Setup(ctx context.Context) error {

	// some setting
	opt := &options.ClientOptions{
		Hosts: []string{"localhost:27017"},
		// Hosts:       []string{"mongodb-primary:27017", "mongodb-secondary:27018", "mongodb-arbiter:27019"},
		MaxPoolSize: refutil.PointUint64(10),
	}

	c, err := mongo.Connect(ctx, opt)

	Client = c
	return err
}

// Ping .
func Ping(ctx context.Context) error {
	return Client.Ping(ctx, readpref.Primary())
}

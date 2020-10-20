package test

import (
	"context"
	"grpc-server/src/infra/mongodb/mconn"
)

// Setup .
func Setup() {
	// mongo
	err := mconn.Setup(context.Background())
	if err != nil {
		panic(err)
	}
}

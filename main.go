package main

import (
	"context"
	"grpc-server/src/infra/grpc"
	"grpc-server/src/infra/mongodb/mconn"
	"log"
)

func main() {
	log.Println("GrpcServer")

	// setup
	err := mconn.Setup(context.Background())
	if err != nil {
		panic(err)
	}

	// server run
	grpc.Run()
}

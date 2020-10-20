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

	// ping
	err = mconn.Ping(context.Background())
	if err != nil {
		log.Println("ping error")
		panic(err)
	}

	// server run
	grpc.Run()
}

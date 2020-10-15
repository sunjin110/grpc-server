package main

import (
	"grpc-server/src/infra/grpc"
	"log"
)

func main() {
	log.Println("GrpcServer")
	grpc.Run()
}

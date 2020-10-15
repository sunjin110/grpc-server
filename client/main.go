package main

import (
	"context"
	"grpc-server/src/infra/grpc/proto/hello"
	"log"

	"google.golang.org/grpc"
)

func main() {
	log.Println("client...")

	conn, err := grpc.Dial("localhost:7766", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	defer conn.Close()

	client := hello.NewHelloClient(conn)
	req := &hello.HelloRequest{Name: "sunjin"}
	res, err := client.Hello(context.TODO(), req)
	if err != nil {
		panic(err)
	}

	log.Println("res message is ", res.Message)
}

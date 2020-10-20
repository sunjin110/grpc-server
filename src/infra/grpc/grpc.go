package grpc

import (
	"grpc-server/src/ctrl/helloctrl"
	"grpc-server/src/infra/grpc/proto/hello"
	"log"
	"net"

	"google.golang.org/grpc"
)

// Run .
func Run() {

	// port
	port, err := net.Listen("tcp", ":7766")
	if err != nil {
		log.Println("portの確保に失敗")
		panic(err)
	}

	// grpc
	server := grpc.NewServer()
	hello.RegisterHelloServer(server, &helloctrl.HelloController{})

	server.Serve(port)
}

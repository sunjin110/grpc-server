package grpc

import (
	"grpc-server/src/ctrl/bankctrl"
	"grpc-server/src/ctrl/bookctrl"
	"grpc-server/src/ctrl/helloctrl"
	"grpc-server/src/infra/grpc/proto/bank"
	"grpc-server/src/infra/grpc/proto/book"
	"grpc-server/src/infra/grpc/proto/hello"

	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
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
	book.RegisterBookRpcServer(server, &bookctrl.BookController{})
	bank.RegisterBankRpcServer(server, &bankctrl.BankController{})

	// grpcurlでアクセスできるようにする
	reflection.Register(server)

	server.Serve(port)
}

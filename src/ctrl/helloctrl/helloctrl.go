package helloctrl

import (
	"context"
	"grpc-server/src/domain/helloserv"
	"grpc-server/src/infra/grpc/proto/hello"
)

// HelloController .
type HelloController struct{}

// Hello .
func (*HelloController) Hello(ctx context.Context, req *hello.HelloRequest) (*hello.HelloReply, error) {
	return helloserv.Hello(ctx, req.Name), nil
}

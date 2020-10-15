package helloctrl

import (
	"context"
	"grpc-server/src/infra/grpc/proto/hello"
)

// HelloController .
type HelloController struct{}

// Hello .
func (*HelloController) Hello(ctx context.Context, req *hello.HelloRequest) (*hello.HelloReply, error) {
	return nil, nil
}

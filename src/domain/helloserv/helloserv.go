package helloserv

import (
	"context"
	"fmt"
	"grpc-server/src/infra/grpc/proto/hello"
)

// Hello 名前を取得して、挨拶を作成する
func Hello(ctx context.Context, name string) *hello.HelloReply {
	if name == "" {
		panic("名前が空白です")
	}

	// some logic ...

	return &hello.HelloReply{
		Message: fmt.Sprintf("こんにちは! %sさん！", name),
	}
}

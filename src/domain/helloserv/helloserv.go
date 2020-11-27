package helloserv

import (
	"context"
	"fmt"
	"grpc-server/src/infra/grpc/proto/hello"
)

// Human 人間
type Human struct {
	Name string
	Age  int32
}

// GetName 人間の名前を
func (human *Human) GetName() string {
	return human.Name
}

// GetName .
func GetName(human Human) string {
	return human.Name
}

// Hello 名前を取得して、挨拶を作成する
func Hello(ctx context.Context, name string) *hello.HelloReply {
	if name == "" {
		panic("名前が空白です")
	}

	sunjin := Human{
		Name: "sunjin",
		Age:  24,
	}

	sunjin.GetName()

	GetName(sunjin)

	// some logic ...

	return &hello.HelloReply{
		Message: fmt.Sprintf("こんにちは! %sさん！", name),
	}
}

package bookctrl

import (
	"context"
	"grpc-server/src/domain/book/bookrpc"
	"grpc-server/src/infra/grpc/proto/book"
	"grpc-server/src/infra/grpc/proto/comm"
)

// BookController .
type BookController struct{}

// Create 本の作成
func (*BookController) Create(ctx context.Context, req *book.CreateRequest) (*comm.Empty, error) {
	bookrpc.Create(ctx, req.Name, req.Price, req.User)
	return &comm.Empty{}, nil
}

// UpdatePrice 本の価格の変更
func (*BookController) UpdatePrice(ctx context.Context, req *book.UpdateRequest) (*comm.Empty, error) {
	bookrpc.UpdatePrice(ctx, req.Name, req.Price)
	return &comm.Empty{}, nil
}

// Delete 本の削除
func (*BookController) Delete(ctx context.Context, req *book.DeleteRequest) (*comm.Empty, error) {
	bookrpc.Delete(ctx, req.Name, req.User)
	return &comm.Empty{}, nil
}

// List リスト
func (*BookController) List(ctx context.Context, req *comm.Empty) (*book.ListReply, error) {
	return bookrpc.List(ctx), nil
}

// DeleteLogList リスト
func (*BookController) DeleteLogList(ctx context.Context, req *comm.Empty) (*book.DeleteLogListReply, error) {
	return bookrpc.DeleteLogList(ctx), nil
}

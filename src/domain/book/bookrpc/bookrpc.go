// Package bookrpc BookRpcService
//
// ControllerからGrpc依存の物を取り除いたロジック(毒抜き的な役割を担う)
// 簡単な実装であれば、ここでロジックを完結してしまっていい
package bookrpc

import (
	"context"
	"grpc-server/src/domain/book/bookcompo"
	"grpc-server/src/infra/grpc/book"
	"grpc-server/src/repo/mongodb/user/bookrepo"
)

// Create 本の作成
func Create(ctx context.Context, name string, price int32) bool {

	// check
	err := bookcompo.CheckBookParam(name, price)
	if err != nil {
		panic(err)
	}

	// create
	bookrepo.Insert(ctx, name, price)
	return true
}

// UpdatePrice 本の価格を変更する
func UpdatePrice(ctx context.Context, name string, price int32) int32 {

	// check
	err := bookcompo.CheckBookParam(name, price)
	if err != nil {
		panic(err)
	}

	// update
	return bookrepo.UpdatePrice(ctx, name, price)
}

// Delete 本の削除
func Delete(ctx context.Context, name string) int32 {
	return bookrepo.Delete(ctx, name)
}

// List 本の一覧を取得
func List(ctx context.Context) *book.ListReply {

	userBookList := bookrepo.GetAll(ctx)
	return bookcompo.CreateListReply(userBookList)
}

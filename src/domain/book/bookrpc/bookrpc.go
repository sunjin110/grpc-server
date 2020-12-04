// Package bookrpc BookRpcService
//
// ControllerからGrpc依存の物を取り除いたロジック(毒抜き的な役割を担う)
// 簡単な実装であれば、ここでロジックを完結してしまっていい
package bookrpc

import (
	"context"
	"grpc-server/src/domain/book/bookcompo"
	"grpc-server/src/infra/grpc/proto/book"
	"grpc-server/src/repo/mongodb/log/booklog"
	"grpc-server/src/repo/mongodb/user/bookrepo"
)

// Create 本の作成
func Create(ctx context.Context, name string, price int32, user string) bool {

	// check
	err := bookcompo.CheckBookParam(name, price)
	if err != nil {
		panic(err)
	}

	// create
	bookrepo.Insert(ctx, name, price, user)
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

// Delete book delete
func Delete(ctx context.Context, name string, user string) bool {

	// check
	err := bookcompo.CheckLogParam(name, user)
	if err != nil {
		panic(err)
	}

	// delete book
	deleteResult := bookrepo.Delete(ctx, name)
	if deleteResult == 0 {
		// error
		panic("本の削除に失敗しました")
	}

	// create
	logResult := booklog.Insert(ctx, name, user)
	if logResult == 0 {
		// error
		panic("削除ログの追加に失敗しました")
	}

	return true
}

// List 本の一覧を取得
func List(ctx context.Context) *book.ListReply {

	userBookList := bookrepo.GetAll(ctx)
	return bookcompo.CreateListReply(userBookList)
}

// BookDeleteHistoryList 削除ログの一覧を取得
func BookDeleteHistoryList(ctx context.Context) *book.BookDeleteHistoryListReply {

	bookDeleteHistoryList := booklog.GetAll(ctx)
	return bookcompo.CreateBookDeleteHistoryListReply(bookDeleteHistoryList)
}

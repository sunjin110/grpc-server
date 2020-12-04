// Package bookcompo BookComponent
//
// コンポーネントとして使いたい部品をここに記述
package bookcompo

import (
	"errors"
	"grpc-server/src/infra/grpc/proto/book"
	"grpc-server/src/repo/mongodb/log/booklog"
	"grpc-server/src/repo/mongodb/user/bookrepo"
)

// CheckBookParam パラメータがおかしくないかをチェックする
func CheckBookParam(name string, price int32) error {

	if name == "" {
		return errors.New("本の名前が空です")
	}

	if price <= 0 {
		return errors.New("本の価格が不正です")
	}

	return nil
}

// CheckLogParam パラメータがおかしくないかをチェックする
func CheckLogParam(name string, user string) error {

	if name == "" {
		return errors.New("本の名前が空です")
	}

	if name == "" {
		return errors.New("ユーザー名が空です")
	}

	return nil
}

// CreateListReply .
func CreateListReply(userBookList []*bookrepo.UserBook) *book.ListReply {

	var bookInfoList []*book.BookInfo
	for _, userBook := range userBookList {

		bookInfo := &book.BookInfo{
			Name:  userBook.Name,
			Price: userBook.Price,
			User:  userBook.User,
		}

		bookInfoList = append(bookInfoList, bookInfo)
	}

	return &book.ListReply{
		BookInfoList: bookInfoList,
	}
}

// CreateBookDeleteHistoryListReply .
func CreateBookDeleteHistoryListReply(bookDeleteHistoryList []*booklog.BookDeleteHistory) *book.BookDeleteHistoryListReply {

	var bookDeleteHistoryInfoList []*book.BookDeleteHistoryInfo
	for _, bookDeleteHistory := range bookDeleteHistoryList {

		bookDeleteHistoryInfo := &book.BookDeleteHistoryInfo{
			Name: bookDeleteHistory.Name,
			User: bookDeleteHistory.User,
		}

		bookDeleteHistoryInfoList = append(bookDeleteHistoryInfoList, bookDeleteHistoryInfo)
	}

	return &book.BookDeleteHistoryListReply{
		BookDeleteHistoryList: bookDeleteHistoryInfoList,
	}
}

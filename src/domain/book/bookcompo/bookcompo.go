// Package bookcompo BookComponent
//
// コンポーネントとして使いたい部品をここに記述
package bookcompo

import "errors"

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

// Package bankrpc BankRpcService
//
// ControllerからGrpc依存の物を取り除いたロジック(毒抜き的な役割を担う)
// 簡単な実装であれば、ここでロジックを完結してしまっていい
package bankrpc

import (
	"context"
	"grpc-server/src/domain/bank/bankcompo"
	"grpc-server/src/infra/grpc/proto/bank"
	"grpc-server/src/repo/mongodb/log/banklog"
	"grpc-server/src/repo/mongodb/user/bankrepo"
)

// CreateBankAccount 口座の作成
func CreateBankAccount(ctx context.Context, user string) bool {

	// check
	err := bankcompo.CheckBankAccountParam(user)
	if err != nil {
		panic(err)
	}

	//名前の重複を確認とか

	// cleate account
	createResult := bankrepo.Insert(ctx, user)
	if createResult == 0 {
		// error
		panic("口座のの作成に失敗しました")
	}

	// 口座作成 を bank_action_history_repo に追加
	logResult := banklog.Insert(ctx, user, "口座開設", 0)
	if logResult == 0 {
		// error
		panic("ログの記録に失敗しました(口座開設)")
	}

	return true
}

// DeleteBankAccount 口座の削除
func DeleteBankAccount(ctx context.Context, user string) bool {

	// check
	err := bankcompo.CheckBankAccountParam(user)
	if err != nil {
		panic(err)
	}

	// delete account
	deleteResult := bankrepo.Delete(ctx, user)
	if deleteResult == 0 {
		// error
		panic("口座の削除に失敗しました")
	}

	// 口座削除 を bank_action_history_repo に追加
	logResult := banklog.Insert(ctx, user, "口座削除", 0)
	if logResult == 0 {
		// error
		panic("ログの記録に失敗しました(口座削除)")
	}

	return true
}

// ListBankAccount 登録されている口座をみる (管理人)
func ListBankAccount(ctx context.Context) *bank.ListBankAccountReply {

	bankAccountList := bankrepo.GetAll(ctx)
	return bankcompo.CreateListBankAccountReply(bankAccountList)
}

// IncUserMoney 入金
func IncUserMoney(ctx context.Context, user string, money int32) bool {

	// check
	err := bankcompo.CheckIncDecMoneyParam(user, money)
	if err != nil {
		panic(err)
	}

	// IncUserMoney
	incMoneyResult := bankrepo.IncUserMoney(ctx, user, money)
	if incMoneyResult == 0 {
		// error
		panic("入金に失敗しました")
	}

	// IncUserMoney を bank_action_history_repo に追加
	logResult := banklog.Insert(ctx, user, "入金", money)
	if logResult == 0 {
		// error
		panic("ログの記録に失敗しました(入金)")
	}
	return true
}

// DecUserMoney 出金
func DecUserMoney(ctx context.Context, user string, money int32) bool {

	// check
	err := bankcompo.CheckIncDecMoneyParam(user, money)
	if err != nil {
		panic(err)
	}

	// DecUserMoney
	decMoneyResult := bankrepo.DecUserMoney(ctx, user, money)
	if decMoneyResult == 0 {
		// error
		panic("出金に失敗しました")
	}

	// IncUserMoney を bank_action_history_repo に追加
	logResult := banklog.Insert(ctx, user, "出金", money)
	if logResult == 0 {
		// error
		panic("ログの記録に失敗しました(出金)")
	}
	return true
}

// TransferUserMoney 送金
func TransferUserMoney(ctx context.Context, userfrom string, userto string, money int32) bool {
	// check
	err := bankcompo.CheckTransferUserMoneyParam(userfrom, userto, money)
	if err != nil {
		panic(err)
	}

	// transfer money
	TransferResult := bankrepo.TransferUserMoney(ctx, userfrom, userto, money)
	if TransferResult == 0 {
		// error
		panic("送金に失敗しました")
	}

	// Transfer を bank_action_history_repo に追加
	logResult := banklog.Insert(ctx, userfrom, "送金", money)
	if logResult == 0 {
		// error
		panic("ログの記録に失敗しました(送金)")
	}
	logResult = banklog.Insert(ctx, userfrom, "受領", money)
	if logResult == 0 {
		// error
		panic("ログの記録に失敗しました(受領)")
	}

	return true

}

// LookUserMoney 残高を確認
func LookUserMoney(ctx context.Context, user string) *bank.LookUserMoneyReply {

	bankAccount := bankrepo.Get(ctx, user)
	return bankcompo.CreateLookUserMoneyReply(bankAccount)
}

// ListUserCashFlow 活動記録を確認
func ListUserCashFlow(ctx context.Context) *bank.ListUserCashFlowReply {

	userCashFlowList := banklog.GetAll(ctx)
	return bankcompo.CreateListUserCashFlowReply(userCashFlowList)
}

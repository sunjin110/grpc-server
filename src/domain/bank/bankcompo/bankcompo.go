// Package bankcompo bankComponent
//
// コンポーネントとして使いたい部品をここに記述
package bankcompo

import (
	"errors"
	"grpc-server/src/infra/grpc/proto/bank"
	"grpc-server/src/repo/mongodb/log/banklog"
	"grpc-server/src/repo/mongodb/user/bankrepo"
)

//// CheckParam
// CheckBankAccountParam
// CheckIncDecMoneyParam
// CheckTransferUserMoneyParam

// CheckBankAccountParam .
func CheckBankAccountParam(user string) error {

	if user == "" {
		return errors.New("ユーザー名が空です")
	}

	return nil
}

// CheckIncDecMoneyParam .
func CheckIncDecMoneyParam(user string, money int32) error {

	if user == "" {
		return errors.New("ユーザー名が空です")
	}

	if money <= 0 {
		return errors.New("金額が不正です")
	}

	return nil
}

// CheckTransferUserMoneyParam .
func CheckTransferUserMoneyParam(userfrom string, userto string, money int32) error {

	if userfrom == "" {
		return errors.New("ユーザー名が空です")
	}

	if userto == "" {
		return errors.New("送金相手の名前が空です")
	}

	if money <= 0 {
		return errors.New("金額が不正です")
	}

	return nil
}

//// Reply
// CreateListBankAccountReply
// CreateListUsersCashFlowReply

// CreateListBankAccountReply .
func CreateListBankAccountReply(bankAccountList []*bankrepo.BankAccount) *bank.ListBankAccountReply {

	var bankAccountInfoList []*bank.BankAccountInfo
	for _, bankAccount := range bankAccountList {

		bankAccountInfo := &bank.BankAccountInfo{
			User:  bankAccount.User,
			Money: bankAccount.Money,
		}

		bankAccountInfoList = append(bankAccountInfoList, bankAccountInfo)
	}

	return &bank.ListBankAccountReply{
		BankAccountInfoList: bankAccountInfoList,
	}
}

// CreateLookUserMoneyReply .
func CreateLookUserMoneyReply(bankAccount *bankrepo.BankAccount) *bank.LookUserMoneyReply {

	bankAccountInfo := &bank.BankAccountInfo{
		User:  bankAccount.User,
		Money: bankAccount.Money,
	}

	return &bank.LookUserMoneyReply{
		BankAccountInfo: bankAccountInfo,
	}
}

// CreateListUserCashFlowReply .
func CreateListUserCashFlowReply(actionList []*banklog.UserCashFlow) *bank.ListUserCashFlowReply {

	var userCashFlowInfoList []*bank.UserCashFlowInfo
	for _, userCashFlow := range actionList {

		actionInfo := &bank.UserCashFlowInfo{
			User:   userCashFlow.User,
			Action: userCashFlow.Action,
			Money:  userCashFlow.Money,
		}

		userCashFlowInfoList = append(userCashFlowInfoList, actionInfo)
	}

	return &bank.ListUserCashFlowReply{
		UserCashFlowInfoList: userCashFlowInfoList,
	}
}

package bankctrl

import (
	"context"
	"grpc-server/src/domain/bank/bankrpc"
	"grpc-server/src/infra/grpc/proto/bank"
	"grpc-server/src/infra/grpc/proto/comm"
)

// BankController .
type BankController struct{}

// CreateBankAccount 口座の作成
func (*BankController) CreateBankAccount(ctx context.Context, req *bank.BankAccountRequest) (*comm.Empty, error) {
	bankrpc.CreateBankAccount(ctx, req.User)
	return &comm.Empty{}, nil
}

// DeleteBankAccount 口座の削除
func (*BankController) DeleteBankAccount(ctx context.Context, req *bank.BankAccountRequest) (*comm.Empty, error) {
	bankrpc.DeleteBankAccount(ctx, req.User)
	return &comm.Empty{}, nil
}

// ListBankAccount 登録されている口座をみる (管理人)
func (*BankController) ListBankAccount(ctx context.Context, req *comm.Empty) (*bank.ListBankAccountReply, error) {
	return bankrpc.ListBankAccount(ctx), nil
}

// IncUserMoney 入金
func (*BankController) IncUserMoney(ctx context.Context, req *bank.UserMoneyRequest) (*comm.Empty, error) {
	bankrpc.IncUserMoney(ctx, req.User, req.Money)
	return &comm.Empty{}, nil
}

// DecUserMoney 出金
func (*BankController) DecUserMoney(ctx context.Context, req *bank.UserMoneyRequest) (*comm.Empty, error) {
	bankrpc.DecUserMoney(ctx, req.User, req.Money)
	return &comm.Empty{}, nil
}

// TranferUserMoney 送金
func (*BankController) TranferUserMoney(ctx context.Context, req *bank.TranferUserMoneyRequest) (*comm.Empty, error) {
	bankrpc.TransferUserMoney(ctx, req.Userfrom, req.Userto, req.Money)
	return &comm.Empty{}, nil
}

// LookUserMoney 個人の残高を確認
func (*BankController) LookUserMoney(ctx context.Context, req *bank.BankAccountRequest) (*bank.LookUserMoneyReply, error) {
	return bankrpc.LookUserMoney(ctx, req.User), nil
}

// ListUserCashFlow 全体の活動記録を確認
func (*BankController) ListUserCashFlow(ctx context.Context, req *comm.Empty) (*bank.ListUserCashFlowReply, error) {
	return bankrpc.ListUserCashFlow(ctx), nil
}

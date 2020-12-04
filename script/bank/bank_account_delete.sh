#!/bin/sh

grpcurl -plaintext -d '{"user":"太郎"}' localhost:7766 bank.BankRpc.DeleteBankAccount
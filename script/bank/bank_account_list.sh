#!/bin/sh

grpcurl -plaintext -d '{}' localhost:7766 bank.BankRpc.ListBankAccount
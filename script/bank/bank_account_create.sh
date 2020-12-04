#!/bin/sh

grpcurl -plaintext -d '{"user":"ひだり"}' localhost:7766 bank.BankRpc.CreateBankAccount
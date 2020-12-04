#!/bin/sh

grpcurl -plaintext -d '{"user":"ひだり", "money":"3000"}' localhost:7766 bank.BankRpc.IncUserMoney
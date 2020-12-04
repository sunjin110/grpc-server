#!/bin/sh

grpcurl -plaintext -d '{"userfrom":"ひだり", "userto":"ひだり", "money":"500"}' localhost:7766 bank.BankRpc.TransferUserMoney
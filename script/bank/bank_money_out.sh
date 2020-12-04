

#!/bin/sh

grpcurl -plaintext -d '{"user":"ひだり", "money":"1000"}' localhost:7766 bank.BankRpc.decUserMoney
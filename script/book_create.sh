#!/bin/sh

grpcurl -plaintext -d '{"name":"ジャンプ", "price":"320"}' localhost:7766 book.BookRpc.Create
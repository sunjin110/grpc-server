#!/bin/sh

grpcurl -plaintext -d '{"name":"ジャンプ", "price":"340", "user":"ひだり"}' localhost:7766 book.BookRpc.Create
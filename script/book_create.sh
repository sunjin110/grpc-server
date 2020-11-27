#!/bin/sh

grpcurl -plaintext -d '{"name":"ジャンプ", "price":"340", "user":"太郎"}' localhost:7766 book.BookRpc.Create
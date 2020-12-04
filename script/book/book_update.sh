#!/bin/sh

grpcurl -plaintext -d '{"name":"ジャンプ", "price":"520"}' localhost:7766 book.BookRpc.UpdatePrice
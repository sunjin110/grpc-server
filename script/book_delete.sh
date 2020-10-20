#!/bin/sh

grpcurl -plaintext -d '{"name":"ジャンプ"}' localhost:7766 book.BookRpc.Delete
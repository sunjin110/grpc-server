#!/bin/sh

grpcurl -plaintext -d '{"name":"ジャンプ", "user":"管理人"}' localhost:7766 book.BookRpc.Delete
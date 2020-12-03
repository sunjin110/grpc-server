#!/bin/sh

grpcurl -plaintext -d '{"name":"サンデー", "price":"340", "user":"ひだり"}' localhost:7766 book.BookRpc.Create
#!/bin/sh

grpcurl -plaintext -d '{}' localhost:7766 book.BookRpc.BookDeleteHistoryList
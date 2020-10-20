#!/bin/sh

# hello
grpcurl -plaintext -d '{"name":"sunjin"}' localhost:7766 hello.Hello.Hello
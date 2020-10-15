#!/bin/sh -ex

DIR=`pwd`

cd ..
OUT_DIR=`pwd`

# clear
rm -rf $DIR/src/infra/grpc/proto
mkdir -p $DIR/src/infra/grpc/proto

export PATH="$PATH:$HOME/go/bin"

cd $DIR/proto
find . -name "*.proto" -exec protoc --proto_path . --go_out=plugins=grpc:$OUT_DIR {} \;

echo "==== create proto fin ===="
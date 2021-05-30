#!/bin/bash

root=$(pwd)
for dir in $root/*; do
    echo $dir
    cd $dir
    protoc --proto_path=$GOPATH/src/heegproto:. --micro_out=. --go_out=. *.proto
done
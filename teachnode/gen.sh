#!/bin/bash

rm -rf *.go
thrift -r -gen go *.thrift

rm -rf gen-go/teachnode/teachnode_service-remote
cp -rf gen-go/teachnode/* ./
rm -rf gen-go

thrift -r -gen html *.thrift
rm -rf ../docs/teachnode
mv gen-html ../docs/teachnode
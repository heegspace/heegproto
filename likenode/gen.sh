#!/bin/bash

rm -rf *.go
thrift -r -gen go *.thrift

rm -rf gen-go/likenode/likenode_service-remote
cp -rf gen-go/likenode/* ./
rm -rf gen-go

thrift -r -gen html *.thrift
mv gen-html ../docs/likenode
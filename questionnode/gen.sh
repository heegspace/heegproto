#!/bin/bash

rm -rf *.go
thrift -r -gen go *.thrift

rm -rf gen-go/questionnode/questionnode_service-remote
cp -rf gen-go/questionnode/* ./
rm -rf gen-go

thrift -r -gen html *.thrift
rm -rf ../docs/questionnode
mv gen-html ../docs/questionnode
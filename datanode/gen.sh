#!/bin/bash

rm -rf *.go
thrift -r -gen go *.thrift

rm -rf gen-go/datanode/datanode_service-remote
cp -rf gen-go/datanode/* ./
rm -rf gen-go

thrift -r -gen html *.thrift
rm -rf ../docs/datanode
mv gen-html ../docs/datanode
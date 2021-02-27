#!/bin/bash

rm -rf *.go
thrift -r -gen go *.thrift

rm -rf gen-go/usernode/usernode_service-remote
cp -rf gen-go/usernode/* ./
rm -rf gen-go

thrift -r -gen html *.thrift
mv gen-html ../docs/usernode
#!/bin/bash

rm -rf *.go
thrift -r -gen go *.thrift

rm -rf gen-go/searchnode/searchnode_service-remote
cp -rf gen-go/searchnode/* ./
rm -rf gen-go

thrift -r -gen html *.thrift
mv gen-html ../docs/searchnode
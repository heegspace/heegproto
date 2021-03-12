#!/bin/bash

rm -rf *.go
thrift -r -gen go *.thrift

rm -rf gen-go/limitnode/limitnode_service-remote
cp -rf gen-go/limitnode/* ./
rm -rf gen-go

thrift -r -gen html *.thrift
mv gen-html ../docs/limitnode
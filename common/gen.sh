#!/bin/bash

rm -rf *.go
thrift -r -gen go *.thrift

cp -rf gen-go/common/* ./
rm -rf gen-go

thrift -r -gen html *.thrift
rm -rf ../docs/common
mv gen-html ../docs/common
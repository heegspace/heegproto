#!/bin/bash

rm -rf *.go
thrift -r -gen go *.thrift

cp -rf gen-go/rescode/* ./
rm -rf gen-go

thrift -r -gen html *.thrift
rm -rf ../docs/rescode
mv gen-html ../docs/rescode
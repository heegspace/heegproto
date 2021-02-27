#!/bin/bash

rm -rf *.go
thrift -r -gen go *.thrift

rm -rf gen-go/notenode/notenode_service-remote
cp -rf gen-go/notenode/* ./
rm -rf gen-go

thrift -r -gen html *.thrift
mv gen-html ../docs/notenode
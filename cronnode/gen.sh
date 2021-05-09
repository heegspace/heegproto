#!/bin/bash

rm -rf *.go
thrift -r -gen go *.thrift

rm -rf gen-go/cronnode/cronnode_service-remote
cp -rf gen-go/cronnode/* ./
rm -rf gen-go


thrift -r -gen html *.thrift
rm -rf ../docs/cronnode
mv gen-html ../docs/cronnode
#!/bin/bash

rm -rf *.go
thrift -r -gen go *.thrift

rm -rf gen-go/certnode/certnode_service-remote
cp -rf gen-go/certnode/* ./
rm -rf gen-go


thrift -r -gen html *.thrift
rm -rf ../docs/certnode
mv gen-html ../docs/certnode
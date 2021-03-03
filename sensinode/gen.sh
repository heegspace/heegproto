#!/bin/bash

rm -rf *.go
thrift -r -gen go *.thrift

rm -rf gen-go/sensinode/sensinode_service-remote
cp -rf gen-go/sensinode/* ./
rm -rf gen-go

thrift -r -gen html *.thrift
rm -rf ../docs/sensinode
mv gen-html ../docs/sensinode
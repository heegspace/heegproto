#!/bin/bash

rm -rf *.go
thrift -r -gen go *.thrift

rm -rf gen-go/cloudnode/cloudnode_service-remote
cp -rf gen-go/cloudnode/* ./
rm -rf gen-go
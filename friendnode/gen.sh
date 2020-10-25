#!/bin/bash

rm -rf *.go
thrift -r -gen go *.thrift

rm -rf gen-go/friendnode/friendnode_service-remote
cp -rf gen-go/friendnode/* ./
rm -rf gen-go
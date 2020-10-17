#!/bin/bash

rm -rf *.go
thrift -r -gen go *.thrift

rm -rf gen-go/registernode/registernode_service-remote
cp -rf gen-go/registernode/* ./
rm -rf gen-go
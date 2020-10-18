#!/bin/bash

rm -rf *.go
thrift -r -gen go *.thrift

rm -rf gen-go/lognode/lognode_service-remote
cp -rf gen-go/lognode/* ./
rm -rf gen-go
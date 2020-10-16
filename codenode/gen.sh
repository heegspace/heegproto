#!/bin/bash

rm -rf *.go
thrift -r -gen go *.thrift

rm -rf gen-go/codenode/codenode_service-remote
cp -rf gen-go/codenode/* ./
rm -rf gen-go
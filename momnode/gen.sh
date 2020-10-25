#!/bin/bash

rm -rf *.go
thrift -r -gen go *.thrift

rm -rf gen-go/momnode/momnode_service-remote
cp -rf gen-go/momnode/* ./
rm -rf gen-go
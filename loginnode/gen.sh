#!/bin/bash

rm -rf *.go
thrift -r -gen go *.thrift

rm -rf gen-go/loginnode/loginnode_service-remote
cp -rf gen-go/loginnode/* ./
rm -rf gen-go
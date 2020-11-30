#!/bin/bash

rm -rf *.go
thrift -r -gen go *.thrift

rm -rf gen-go/authnode/authnode_service-remote
cp -rf gen-go/authnode/* ./
rm -rf gen-go
#!/bin/bash

rm -rf *.go
thrift -r -gen go *.thrift

rm -rf gen-go/dartynode/dartynode_service-remote
cp -rf gen-go/dartynode/* ./
rm -rf gen-go
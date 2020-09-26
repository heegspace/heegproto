#!/bin/bash

rm -rf *.go
thrift -r -gen go *.thrift

rm -rf gen-go/s2sname/s2sname_service-remote
cp -rf gen-go/s2sname/* ./
rm -rf gen-go
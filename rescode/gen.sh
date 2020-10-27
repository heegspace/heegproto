#!/bin/bash

rm -rf *.go
thrift -r -gen go *.thrift

cp -rf gen-go/rescode/* ./
rm -rf gen-go
#!/bin/bash

rm -rf rescode
thrift -r -gen go rescode.thrift

cp -rf gen-go/heegproto/rescode ./
rm -rf gen-go
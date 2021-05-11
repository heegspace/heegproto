#!/bin/bash

rm -rf *.go
thrift -r -gen go *.thrift

rm -rf gen-go/macipnode/macipnode_service-remote
cp -rf gen-go/macipnode/* ./
rm -rf gen-go


thrift -r -gen html *.thrift
rm -rf ../docs/macipnode
mv gen-html ../docs/macipnode
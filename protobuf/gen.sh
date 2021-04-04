#!/bin/bash

protoc -I. --gofast_out=. *.proto
mv *go ../protocol/
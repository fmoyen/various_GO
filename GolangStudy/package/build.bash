#!/bin/bash

prog=package_example.go
export GOBIN=$PWD/bin  # directory where binary prog will be installed
export GOPATH=$PWD  # package files in $GOPATH/src directory (subdirectory names = package namesi like: $GOPATH/src/calculation/calc.go)

mkdir -p $GOBIN

go install $prog

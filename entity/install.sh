#!/bin/bash

dir=$GOPATH

echo "GOPATH="$dir

go build -o gormEntity ./main.go

mv gormEntity $dir/bin

echo "gormEntity install success"
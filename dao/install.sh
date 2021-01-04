#!/bin/bash

dir=$GOPATH

echo "GOPATH="$dir

go build -o gormDao ./main.go

mv gormDao $dir/bin

echo "gormDao install success"
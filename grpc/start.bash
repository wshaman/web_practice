#!/bin/bash
go get -u github.com/golang/protobuf/protoc-gen-go
go get google.golang.org/grpc
cd $GOPATH/src/github.com/golang/protobuf/protoc-gen-go || exit
go build -o $GOPATH/bin .
protoc -I proto/ proto/server.proto --go_out=plugins=grpc:server

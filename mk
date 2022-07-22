#!/bin/sh

INPUT_FILE=pkg/pb/chat.proto

proto(){
    # Don't forget to
    # go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
    # go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
    export PATH="$PATH:$(go env GOPATH)/bin"

    # Generate Go code
    protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative \
        $INPUT_FILE
}

"$@"
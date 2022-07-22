#!/bin/sh

INPUT_FILE=pkg/pb/chat.proto

proto(){
    export PATH="$PATH:$(go env GOPATH)/bin"

    # Generate Go code
    protoc --go_out=. --go_opt=paths=source_relative \
           --go-grpc_out=. --go-grpc_opt=paths=source_relative \
           $INPUT_FILE
}

"$@"
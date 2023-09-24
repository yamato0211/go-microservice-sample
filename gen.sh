#!/bin/bash
cd user/proto/
protoc --go_out=. --go-grpc_out=. user.proto
cd ../..
cd todo/proto/
protoc --go_out=. --go-grpc_out=. todo.proto
cd ../..

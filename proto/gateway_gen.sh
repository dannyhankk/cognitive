#!/bin/zsh
protoc -I . --grpc-gateway_out . \
    --grpc-gateway_opt logtostderr=true \
    --grpc-gateway_opt paths=source_relative \
    --grpc-gateway_opt generate_unbound_methods=true \
    --grpc-gateway_opt grpc_api_configuration=../text2speech.yaml \
    ./*.proto
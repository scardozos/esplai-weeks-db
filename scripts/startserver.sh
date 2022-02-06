#!/bin/bash
export GRPC_ADDR_SELF="0.0.0.0:9001"
go run -race pkg/main.go
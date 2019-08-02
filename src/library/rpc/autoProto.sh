#!/usr/bin/env bash

protoc -I . grpc.proto --go_out=plugins=grpc:.
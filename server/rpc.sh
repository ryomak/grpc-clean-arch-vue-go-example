#!/bin/bash
protoc -I ./proto user.proto --go_out=plugins=grpc:./
protoc -I ./proto room.proto --go_out=plugins=grpc:./

#!/bin/bash
for file in `\find bff -name '*.proto'`; do
    temp="$temp $file"
done
protoc --go_out=plugins=grpc:. $temp
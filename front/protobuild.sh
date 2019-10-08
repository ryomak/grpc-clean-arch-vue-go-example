protoc -I ./../proto user.proto --js_out=import_style=commonjs:./src/assets --grpc-web_out=import_style=commonjs,mode=grpcwebtext:./src/assets
protoc -I ./../proto room.proto --js_out=import_style=commonjs:./src/assets --grpc-web_out=import_style=commonjs,mode=grpcwebtext:./src/assets

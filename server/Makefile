
proto-build:
	protoc -I ./../proto user.proto --go_out=plugins=grpc:./src/interface/rpc
	protoc -I ./../proto room.proto --go_out=plugins=grpc:./src/interface/rpc

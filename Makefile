build_protos_go:
	if exist ./go_protos/*.go del go_protos\*.go
	protoc --proto_path=protos --go_out=plugins=grpc,import_path=go_protos,import_prefix=go_protos:./go_protos/ protos/*.proto



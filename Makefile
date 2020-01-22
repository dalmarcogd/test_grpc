build_protos_go:
	if not exist ./go_protos mkdir "go_protos"
	if exist ./go_protos/*.go del go_protos\*.go
	protoc --proto_path=protos --go_out=plugins=grpc,import_path=protos:./go_protos/ protos/*.proto
	xcopy "go_protos\\*.*" "go_client\protos\\" /s /h /e /k /f /c /y
	xcopy "go_protos\\*.*" "go_server\protos\\" /s /h /e /k /f /c /y
	echo S|rmdir /s go_protos

build_protos_python:
	if exist ./python_protos/*.py del python_protos\*.py
	protoc --proto_path=protos --python_out=plugins=grpc,import_path=python_protos:./python_protos/ protos/*.proto



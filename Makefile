all:
	protoc -I proto ./proto/tasks/*.proto --go_out=./gen/go/tasks/ --go-grpc_out=./gen/go/tasks/
	protoc -I proto ./proto/users/*.proto --go_out=./gen/go/users/ --go-grpc_out=./gen/go/users/

 
protogen: 	
	protoc -I ./api/proto --go_out=./pkg/investapi --go-grpc_out=./pkg/investapi ./api/proto/*.proto
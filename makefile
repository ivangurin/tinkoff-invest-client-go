protogen: 	
	protoc -I ./api/proto --go_out=./internal/investapi --go-grpc_out=./internal/investapi ./api/proto/*.proto
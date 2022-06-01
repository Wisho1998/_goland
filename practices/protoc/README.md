> protoc proto/entities/*.proto proto/mongo/*.proto --go-grpc_out=../protoc --go_out=../protoc>
> 
> protoc-go-inject-tag -input=pb/*.pb.go
> 
> go run main
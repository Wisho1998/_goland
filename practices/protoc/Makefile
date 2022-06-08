run:
	go run .
gen:
	protoc proto/entities/*.proto proto/mongo/*.proto --go-grpc_out=../protoc --go_out=../protoc
	make gen_tag
gen_tag:
	protoc-go-inject-tag -input=pb/*.pb.go
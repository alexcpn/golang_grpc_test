
CURRENTDIR = /home/alex/coding/golang_grpc_test
all: get_pgv  build_proto_with_validate build_go

# get the protoc validator
# https://github.com/envoyproxy/protoc-gen-validate
get_pgv:
#Note this is using go with modules go version 1.12
	go get  github.com/envoyproxy/protoc-gen-validate@v0.1.0
build_proto:
	protoc -I interfaces/ interfaces/*.proto --go_out=plugins=grpc:./interfaces
build_proto_with_validate:
	protoc   -I interfaces/  -I ${GOPATH}/pkg/mod/github.com/envoyproxy/protoc-gen-validate@v0.1.0/ \
	 --go_out=plugins=grpc:./interfaces   --validate_out=lang=go:./interfaces interfaces/*.proto

build_go:
	cd test_server && go build
	cd test_client && go build
# Build inside docker	
docker-build:
	echo $(CURRENTDIR)
	docker run --rm -it -v $(CURRENTDIR)\:/go namely/protoc-all make all
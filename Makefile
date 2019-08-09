
CURRENTDIR =  $(shell pwd)
HTTP_PROXY = http://10.158.100.6:8080 #change this as per your proxy

all: get_pgv  build_proto_with_validate build_go build_proto_documentation

# get the protoc validator
# https://github.com/envoyproxy/protoc-gen-validate
get_pgv:
# Note this is using go with modules go version 1.12
# To make the build faster these can be added to the base Docker build imagage
	go get  github.com/envoyproxy/protoc-gen-validate@v0.1.0
	go get  github.com/pseudomuto/protoc-gen-doc/cmd/protoc-gen-doc@v1.3.0
build_proto:
	protoc -I interfaces/ interfaces/*.proto --go_out=plugins=grpc:./interfaces
build_proto_with_validate:
	protoc   -I interfaces/  -I ${GOPATH}/pkg/mod/github.com/envoyproxy/protoc-gen-validate@v0.1.0/ \
	 --go_out=plugins=grpc:./interfaces   --validate_out=lang=go:./interfaces interfaces/*.proto
build_proto_documentation:
	protoc   -I interfaces/ -I ${GOPATH}/pkg/mod/github.com/envoyproxy/protoc-gen-validate@v0.1.0/ \
	--plugin=protoc-gen-doc=${GOPATH}/bin/protoc-gen-doc \
     --doc_out=./doc --doc_opt=markdown,index.md interfaces/*.proto
	 #--doc_out=./doc --doc_opt=html,index.html 

build_go:
	cd integration_test && go test
	cd test_server && go build
	cd test_client && go build

# Build a Docker conatiner: if you do this you don't have to run docker-make
docker-build:
	# this is the base container -only needed if THIS Container (grpc/go:1.12) does not exist
	docker build --build-arg  http_proxy=$(HTTP_PROXY) \
	  --build-arg https_proxy=$(HTTP_PROXY)  -t grpc/go:1.12 - < Dockerfile.go_grpc.1.12 
	# next we use the multi-stage build to build a container with the binary
	docker build --build-arg  http_proxy=$(HTTP_PROXY) \
	 --build-arg https_proxy=$(HTTP_PROXY)  -t go_grpc:1.0 -f  Dockerfile $(CURRENTDIR) 


# Build inside docker with volume mapping to test locally
docker-make:
	echo "ENV variables check" $(CURRENTDIR) $(HTTP_PROXY)  

	docker run  --env http_proxy=$(HTTP_PROXY) --env https_proxy=$(HTTP_PROXY) \
	--rm -it -v $(CURRENTDIR):/go/golang_grpc_test grpc/go:1.12 /bin/sh -c \
	 'cd golang_grpc_test && make all'
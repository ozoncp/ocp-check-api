.PHONY: build
build: vendor-proto .generate .build

PHONY: .generate
.generate:
		mkdir -p swagger
		mkdir -p pkg/ocp-check-api
		protoc -I vendor.protogen \
				--go_out=pkg/ocp-check-api --go_opt=paths=import \
				--go-grpc_out=pkg/ocp-check-api --go-grpc_opt=paths=import \
				--grpc-gateway_out=pkg/ocp-check-api \
				--grpc-gateway_opt=logtostderr=true \
				--grpc-gateway_opt=paths=import \
				--validate_out lang=go:pkg/ocp-check-api \
				--swagger_out=allow_merge=true,merge_file_name=check:swagger \
				api/ocp-check-api/ocp-check-api.proto
		mv pkg/ocp-check-api/github.com/ozoncp/ocp-check-api/pkg/ocp-check-api/* pkg/ocp-check-api/
		rm -rf pkg/ocp-check-api/github.com
		mkdir -p cmd/ocp-check-api
		mkdir -p pkg/ocp-test-api
		protoc -I vendor.protogen \
				--go_out=pkg/ocp-test-api --go_opt=paths=import \
				--go-grpc_out=pkg/ocp-test-api --go-grpc_opt=paths=import \
				--grpc-gateway_out=pkg/ocp-test-api \
				--grpc-gateway_opt=logtostderr=true \
				--grpc-gateway_opt=paths=import \
				--validate_out lang=go:pkg/ocp-test-api \
				--swagger_out=allow_merge=true,merge_file_name=test:swagger \
				api/ocp-test-api/ocp-test-api.proto
		mv pkg/ocp-test-api/github.com/ozoncp/ocp-check-api/pkg/ocp-test-api/* pkg/ocp-test-api/
		rm -rf pkg/ocp-test-api/github.com

PHONY: .build
.build:
		CGO_ENABLED=0 GOOS=linux go build -o bin/ocp-check-api cmd/ocp-check-api/main.go

PHONY: install
install: build .install

PHONY: .install
install:
		go install cmd/grpc-server/main.go

PHONY: vendor-proto
vendor-proto: .vendor-proto

PHONY: .vendor-proto
.vendor-proto:
		mkdir -p vendor.protogen
		mkdir -p vendor.protogen/api/ocp-check-api
		cp api/ocp-check-api/ocp-check-api.proto vendor.protogen/api/ocp-check-api
		@if [ ! -d vendor.protogen/google ]; then \
			git clone https://github.com/googleapis/googleapis vendor.protogen/googleapis &&\
			mkdir -p  vendor.protogen/google/ &&\
			mv vendor.protogen/googleapis/google/api vendor.protogen/google &&\
			rm -rf vendor.protogen/googleapis ;\
		fi
		@if [ ! -d vendor.protogen/github.com/envoyproxy ]; then \
			mkdir -p vendor.protogen/github.com/envoyproxy &&\
			git clone https://github.com/envoyproxy/protoc-gen-validate vendor.protogen/github.com/envoyproxy/protoc-gen-validate ;\
		fi
		mkdir -p vendor.protogen/api/ocp-test-api
		cp api/ocp-test-api/ocp-test-api.proto vendor.protogen/api/ocp-test-api
		@if [ ! -d vendor.protogen/google ]; then \
			git clone https://github.com/googleapis/googleapis vendor.protogen/googleapis &&\
			mkdir -p  vendor.protogen/google/ &&\
			mv vendor.protogen/googleapis/google/api vendor.protogen/google &&\
			rm -rf vendor.protogen/googleapis ;\
		fi
		@if [ ! -d vendor.protogen/github.com/envoyproxy ]; then \
			mkdir -p vendor.protogen/github.com/envoyproxy &&\
			git clone https://github.com/envoyproxy/protoc-gen-validate vendor.protogen/github.com/envoyproxy/protoc-gen-validate ;\
		fi

.PHONY: deps
deps: install-go-deps

.PHONY: install-go-deps
install-go-deps: .install-go-deps

.PHONY: .install-go-deps
.install-go-deps:
		ls go.mod || go mod init
		go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
		go get -u github.com/golang/protobuf/proto
		go get -u github.com/golang/protobuf/protoc-gen-go
		go install github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
		go install google.golang.org/grpc/cmd/protoc-gen-go-grpc
		go install github.com/envoyproxy/protoc-gen-validate

.PHONY: install-migrate
install-migrate: .install-migrate

.PHONY: .install-migrate
.install-migrate:
		curl -L https://github.com/golang-migrate/migrate/releases/download/v4.14.1/migrate.linux-amd64.tar.gz | tar xvz && mv ./migrate.linux-amd64 ./migrate
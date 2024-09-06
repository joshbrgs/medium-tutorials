# Paths
PROTO_DIR := ./cmd/nemesis/api/proto
GEN_DIR := ./cmd/nemesis/api/gen

# Commands
PROTOC := protoc
PROTOC_GEN_GO := $(shell go env GOPATH)/bin/protoc-gen-go
PROTOC_GEN_GO_GRPC := $(shell go env GOPATH)/bin/protoc-gen-go-grpc

# Proto files
PROTO_FILES := $(wildcard $(PROTO_DIR)/*.proto)

# Ensure protoc-gen-go and protoc-gen-go-grpc are available
$(PROTOC_GEN_GO):
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest

$(PROTOC_GEN_GO_GRPC):
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

# Rule to generate Go code from proto files
.PHONY: proto
proto: $(PROTOC_GEN_GO) $(PROTOC_GEN_GO_GRPC)
	$(PROTOC) \
		--proto_path=$(PROTO_DIR) \
		--go_out=$(GEN_DIR) \
		--go_opt=paths=source_relative \
		--go-grpc_out=$(GEN_DIR) \
		--go-grpc_opt=paths=source_relative \
		$(PROTO_FILES)

# Cleanup generated files
.PHONY: clean
clean:
	rm -rf $(GEN_DIR)

# Default target
.PHONY: all
all: proto

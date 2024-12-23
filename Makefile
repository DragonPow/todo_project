proto:
	protoc \
		-I ./third_party/ \
		--go_out . --go-grpc_out . --grpc-gateway_out . \
		--proto_path $(MODULE_DIR)/api $(MODULE_DIR)/api/*.proto

update:
	go mod tidy
	go mod vendor
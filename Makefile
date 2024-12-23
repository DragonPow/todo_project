proto:
	protoc \
		-I ./third_party/ \
		--go_out . --go-grpc_out . --grpc-gateway_out . \
		--proto_path $(MODULE_DIR)/api $(MODULE_DIR)/api/*.proto

update:
	go mod tidy
	go mod vendor

# Build file

build-todo:
	docker build -t todo-web -f /app/todo/Dockerfile .

build-iam:
	docker build -t iam-web -f /app/todo/Dockerfile .

# Run file

run-todo:
	docker run -p 8080:8080 todo-web

run-iam:
	docker run -p 8081:8081 iam-web

run-all:
	docker compose up --build -d
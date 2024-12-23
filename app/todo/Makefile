# Add api dir from root project
MODULE_DIR = ./app/todo

proto:
	cd ../.. && make proto MODULE_DIR=$(MODULE_DIR)

migrate-up:
	go run ./cmd/*.go migrate-up

run:
	go run ./cmd/*.go server

update:
	cd ../.. && make update
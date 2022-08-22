say_hello:
	@echo "Hello Golang!!!!"

install:
	@echo "Build a go binary"
	@cd cmd/wedding_gifts && go build -o bin/wedding_gifts
	@./cmd/wedding_gifts/bin/wedding_gifts
run:
	@make build
	./random-quotes

build:
	@go build -v -o random-quotes ./cmd/main.go
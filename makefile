build:
	go build -o bin/currency

run: build
	bin/currency
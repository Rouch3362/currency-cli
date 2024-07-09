build:
	go build -o bin/currency main.go

run: build
	bin/currency
install:
	dep ensure

build:
	go build -o main

start:
	./main
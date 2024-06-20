SHELL = bash

all: proto build

proto:
	cd slowpizza && protoc --go_out=../. --go-grpc_out=../. ./slowpizza.proto

build: build-client build-server

build-client:
	go build -o slowpizza-client ./client

build-server:
	go build -o slowpizza-server ./server

image:
	docker build . -t slowpizza:latest

all: setup
.PHONY: all

setup:
	go-bindata -debug -o ./server/bindata.go assets/*

build:
	cross-env NODE_ENV=production webpack -p --config ./client/webpack/webpack.config.js && go-bindata -o ./server/bindata.go assets/* && gox -osarch=\"windows/amd64 linux/amd64\" -output=dist/web_{{.OS}}_{{.Arch}} ./server/*
.PHONY: build

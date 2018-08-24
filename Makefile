all: setup
.PHONY: all

setup:
	go-bindata -debug -o ./server/bindata.go assets/*

build:
	go-bindata -debug -o ./server/bindata.go assets/*
	go build -tags=release -o bin/gin  server/main.go server/router.go server/bindata.go
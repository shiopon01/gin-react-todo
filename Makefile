all: setup
.PHONY: all

setup:
	go-bindata -debug -o ./server/bindata.go assets/*

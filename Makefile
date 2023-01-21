set-env:
	cp .envrc.example .envrc

deps:
	go get

build:
	go build .

run:
	go run .

vet:
	go vet .

setup: set-env deps build

all: setup run
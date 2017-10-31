.PHONY: all generate build link
all: generate build link

generate:
	go generate

build: generate
	go build -o terraform-provider-contrail

link:
	cd examples && ln -sf ../terraform-provider-contrail


.PHONY: build
name = smaker

build:
	go build -ldflags "-X main._VERSION_=$(shell date +%Y%m%d-%H%M%S)" -o $(name)
	mv $(name) bin/$(name)

run: build
	bin/$(name)

install: build
	cp bin/$(name) /usr/local/bin

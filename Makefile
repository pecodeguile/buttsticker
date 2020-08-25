SHELL := /bin/sh

GOOS ?= linux
GOARCH ?= amd64

.PHONY: all clean

all: buttstickerapi

clean:
	$(RM) buttstickerapi

buttstickerapi: cmd/buttstickerapi/buttstickerapi.go
	CGO_ENABLED=0 GOOS=$(GOOS) GOARCH=$(GOARCH) go build -a -ldflags '-extldflags "-static"' -o $@ $^

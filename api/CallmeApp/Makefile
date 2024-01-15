# Makefile for CallmeApplication

# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOTEST=$(GOCMD) test
GOLINT=golint
GOVET=$(GOCMD) vet
GOFMT=gofmt -s -w

# Go files
GOFILES=CallmeApplication.go CallmeController.go

.DEFAULT_GOAL := all

all: fmt lint vet test build

build: 
	$(GOBUILD) $(GOFILES)
.PHONY:build

test: 
	$(GOTEST) -v ./...

lint: 
	$(GOLINT) $(GOFILES)
.PHONY:lint

vet:
	$(GOVET) $(GOFILES)
.PHONY:vet

fmt:
	$(GOFMT) $(GOFILES)
.PHONY:fmt

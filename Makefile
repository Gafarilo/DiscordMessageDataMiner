
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
BINARY_NAME=DiscordScraper
VERSION=$(shell cat ./VERSION)
BUILD_TIME=$(shell date -u +"%Y-%m-%dT%H:%M:%SZ")
COMMIT=$(shell git rev-parse --short HEAD)

LDFLAGS=-ldflags "-X=main.version=$(VERSION) -X=main.buildTime=$(BUILD_TIME) -X main.gitCommit=$(COMMIT)"

all: build
SHELL := /bin/bash

.PHONY: tools
tools:
	go install github.com/vektra/mockery/v2@latest

.PHONY: mocks
mocks:
	mockery --all

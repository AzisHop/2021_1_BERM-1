#!/bin/bash
golangci-lint run -E lll -E prealloc -E unconvert -E unparam -E whitespace -E gomnd -E testpackage  ./...

#!/bin/bash
golangci-lint run -E golint -E lll -E prealloc -E unconvert -E unparam -E whitespace -E gomnd -E testpackage  -D deadcode -D varcheck -D unused -D structcheck

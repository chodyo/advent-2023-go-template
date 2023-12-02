# Advent of Code 2023: Go Template

[![Go](https://github.com/chodyo/advent-2023-go-template/actions/workflows/go.yml/badge.svg)](https://github.com/chodyo/advent-2023-go-template/actions/workflows/go.yml)

## Setup Instructions

1. Create a new repo from this template.
2. Uncomment the "expected" test values in `internal/commands/<commandName>/<commandName>_test.go` and remove the "expect 0" lines.
3. Implement your algorithm in `internal/commands/<commandName>/<commandName>.go`.

## Running

Using `go run . -h` will print help text that will guide you on how to run this code.

## Testing

`go test -v ./...`

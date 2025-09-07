# Go Concurrency Tasks Makefile

.PHONY: help test test-all test-coverage test-race clean install

# Default target
help:
	@echo "Available commands:"
	@echo "  make test          - Run all tests"
	@echo "  make test-all      - Run all tests with verbose output"
	@echo "  make test-coverage - Run tests with coverage report"
	@echo "  make test-race     - Run tests with race detection"
	@echo "  make test-task1    - Run tests for task 1 (goroutines)"
	@echo "  make test-task2    - Run tests for task 2 (channels)"
	@echo "  make test-task3    - Run tests for task 3 (mutex)"
	@echo "  make test-task4    - Run tests for task 4 (select)"
	@echo "  make test-task5    - Run tests for task 5 (worker pool)"
	@echo "  make test-task6    - Run tests for task 6 (fan-out fan-in)"
	@echo "  make test-task7    - Run tests for task 7 (context)"
	@echo "  make test-task8    - Run tests for task 8 (pipeline)"
	@echo "  make test-task9    - Run tests for task 9 (rate limiter)"
	@echo "  make test-task10   - Run tests for task 10 (barrier)"
	@echo "  make test-task11   - Run tests for task 11 (semaphore)"
	@echo "  make test-task12   - Run tests for task 12 (race condition)"
	@echo "  make test-task13   - Run tests for task 13 (circuit breaker)"
	@echo "  make test-task14   - Run tests for task 14 (bulkhead)"
	@echo "  make test-task15   - Run tests for task 15 (retry)"
	@echo "  make test-task16   - Run tests for task 16 (timeout)"
	@echo "  make test-task17   - Run tests for task 17 (broadcast)"
	@echo "  make test-task18   - Run tests for task 18 (priority queue)"
	@echo "  make test-task19   - Run tests for task 19 (cache)"
	@echo "  make test-task20   - Run tests for task 20 (actor model)"
	@echo "  make test-task21   - Run tests for task 21 (event sourcing)"
	@echo "  make test-task22   - Run tests for task 22 (stream processing)"
	@echo "  make install       - Install dependencies"
	@echo "  make clean         - Clean test cache"

# Install dependencies
install:
	go mod tidy
	go mod download

# Run all tests
test:
	go test ./...

# Run all tests with verbose output
test-all:
	go test -v ./...

# Run tests with coverage
test-coverage:
	go test -cover ./...
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report generated: coverage.html"



# Run tests with race detection
test-race:
	go test -race ./...

# Individual task tests
test-task1:


	go test -v ./task1

test-task2:
	go test -v ./task2

test-task3:
	go test -v ./task3

test-task4:
	go test -v ./task4

test-task5:
	go test -v ./task5

test-task6:
	go test -v ./task6

test-task7:
	go test -v ./task7

test-task8:
	go test -v ./task8

test-task9:
	go test -v ./task9

test-task10:
	go test -v ./task10

test-task11:
	go test -v ./task11

test-task12:
	go test -v ./task12

test-task13:
	go test -v ./task13

test-task14:
	go test -v ./task14

test-task15:
	go test -v ./task15

test-task16:
	go test -v ./task16

test-task17:
	go test -v ./task17

test-task18:
	go test -v ./task18

test-task19:
	go test -v ./task19

test-task20:
	go test -v ./task20

test-task21:
	go test -v ./task21

test-task22:
	go test -v ./task22

# Clean test cache
clean:
	go clean -testcache
	rm -f coverage.out coverage.html

# Run specific test with race detection
test-race-task1:
	go test -race ./task1

test-race-task2:
	go test -race ./task2

test-race-task3:
	go test -race ./task3

test-race-task4:
	go test -race ./task4

test-race-task5:
	go test -race ./task5

test-race-task6:
	go test -race ./task6

test-race-task7:
	go test -race ./task7

test-race-task8:
	go test -race ./task8

test-race-task9:
	go test -race ./task9

test-race-task10:
	go test -race ./task10

test-race-task11:
	go test -race ./task11

test-race-task12:
	go test -race ./task12

test-race-task13:
	go test -race ./task13

test-race-task14:
	go test -race ./task14

test-race-task15:
	go test -race ./task15

test-race-task16:
	go test -race ./task16

test-race-task17:
	go test -race ./task17

test-race-task18:
	go test -race ./task18

test-race-task19:
	go test -race ./task19

test-race-task20:
	go test -race ./task20

test-race-task21:
	go test -race ./task21

test-race-task22:
	go test -race ./task22

# Benchmark tests (if implemented)
bench:
	go test -bench=. ./...

# Format code
fmt:
	go fmt ./...

# Lint code
lint:
	golangci-lint run

# Build all tasks
build:
	go build ./...

# Run all tasks (if they have main functions)
run:
	@echo "Tasks are designed to be tested, not run directly"
	@echo "Use 'make test' to verify your implementations"

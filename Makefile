PHONY: test-coverage
test-coverage:
	@echo "Running tests..."
	go test -cover ./... -coverprofile=c.out
	go tool cover -html=c.out -o coverage.html
	rm -rf c.out
	open coverage.html

.PHONY: test
test:
	@echo "Running tests..."
	go test ./... -v

.PHONY: run
run:
	@echo "Starting application..."
	go run src/cmd/cli/main.go


.PHONY: fmt lint

fmt:
	golangci-lint fmt

lint:
	golangci-lint custom -v
	./custom-gcl run --fix ./...
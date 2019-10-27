all: gotool
	@go build -o main -v .
gotool:
	gofmt -w .
	go tool vet . | grep -v vendor;true
test:
	@go test -v -count=1  ./...

help:
	@echo "make - compile the source code with local vendor"
	@echo "make build compile the source code with remote vendor"
	@echo "make gotool - run go tool 'fmt' and 'vet'"

.PHONY: clean gotool ca help

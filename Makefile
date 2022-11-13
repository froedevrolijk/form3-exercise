.PHONY: lint
lint:
	@echo "--> Running golangci"
	golangci-lint run ./form3

.PHONY: fmt
fmt:
	@echo "--> Running go fmt"
	go fmt ./form3

.PHONY: test
test:
	@echo "--> Running tests"
	go test -v -coverprofile=coverage.out -covermode=atomic ./form3

.PHONY: docker-test
docker-test:
	docker-compose up --build --abort-on-container-exit apiclient
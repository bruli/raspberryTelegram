define help
Usage: make <command>
Commands:
   help:                      Show this help information
   test:                      Run unit tests
   build:                     Compile the project
   coverage:                  Run unit tests with coverage
   encryptVault:              Encrypt vault secret file
   decryptVault:              Decrypt vault secret file
   deploy:                    Deploy the code to raspberry
   lint:                      Execute go linter
   docker-exec-builder:       Start builder docker container and entry inside it. Build project here.
endef
export help

.PHONY: help
help:
	@echo "$$help"

.PHONY: test
test:
	go test -race ./...

.PHONY: build
build:
	CGO_ENABLED=0 GOOS=linux GOARCH=arm GOARM=6 go build -o devops/ansible/assets/telegram cmd/telegram/main.go

.PHONY: coverage
coverage:
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out

.PHONY: encryptVault
encryptVault:
	ansible-vault encrypt --vault-id raspberry_telegram@devops/ansible/password devops/ansible/inventories/production/group_vars/raspberry_telegram/vault.yml

.PHONY: decryptVault
decryptVault:
	ansible-vault decrypt --vault-id raspberry_telegram@devops/ansible/password devops/ansible/inventories/production/group_vars/raspberry_telegram/vault.yml

.PHONY: deploy
deploy:
	devops/scripts/deploy.sh

.PHONY: fumpt
fumpt:
	go tool gofumpt -w -l .

.PHONY: lint
lint:
	go tool golangci-lint run
	go mod tidy -v && git --no-pager diff --quiet go.mod go.sum

.PHONY: docker-exec-builder
docker-exec-builder:
	docker build -t builder .
	docker run -it --rm -v $(shell pwd):/app builder bash
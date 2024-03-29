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

test:
	go test -race ./...

build:
	CC=arm-linux-gnueabi-gcc CGO_ENABLED=1 GOOS=linux GOARCH=arm GOARM=6 go build -o devops/ansible/assets/telegram cmd/telegram/main.go

coverage:
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out

encryptVault:
	ansible-vault encrypt --vault-id raspberry_telegram@devops/ansible/password devops/ansible/inventories/production/group_vars/raspberry_telegram/vault.yml
decryptVault:
	ansible-vault decrypt --vault-id raspberry_telegram@devops/ansible/password devops/ansible/inventories/production/group_vars/raspberry_telegram/vault.yml

deploy:
	devops/scripts/deploy.sh

lint:
	golangci-lint run
	go mod tidy -v && git --no-pager diff --quiet go.mod go.sum

docker-exec-builder:
	docker build -t builder .
	docker run -it --rm -v $(shell pwd):/app builder bash
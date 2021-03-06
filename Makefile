tests:
	go test ./...

acceptance:
	bash -c "cd test/acceptance && go test -v"

unit:
	@bash -c "cd internal && GOTEST_PALETTE="red,blue" gotest ./..."

build:
	CC=arm-linux-gnueabi-gcc CGO_ENABLED=1 GOOS=linux GOARCH=arm GOARM=6 go build -o deploy/assets/telegram cmd/telegram/main.go

deploy:
	@make build
	cd cmd/telegram && scp telegram waterSystem:

coverage:
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out

encryptVault:
	ansible-vault encrypt --vault-id raspberry_telegram@deploy/password deploy/inventories/production/group_vars/raspberry_telegram/vault.yml
decryptVault:
	ansible-vault decrypt --vault-id raspberry_telegram@deploy/password deploy/inventories/production/group_vars/raspberry_telegram/vault.yml

deploys:
	ansible-playbook -i deploy/inventories/production/hosts deploy/deploy.yml --vault-id raspberry_telegram@deploy/password
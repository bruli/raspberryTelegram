coverage:
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out

tests:
	go test ./...

build:
	cd cmd/telegram && CC=arm-linux-gnueabi-gcc CGO_ENABLED=1 GOOS=linux GOARCH=arm GOARM=6 go build

deploy:
	@make build
	cd cmd/telegram && scp telegram waterSystem:

coverage:
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out
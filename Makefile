run: test
	clear && go run main.go

test:
	go test -v ./...
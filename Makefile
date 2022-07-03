run: test
	clear && go run main.go

test:
	go test -v ./...

mockgen:
	@go generate ./...

start-mqtt:
	docker run -it -p 1883:1883 -p 9001:9001 -v /Users/farcun/Works/Xendit/escrow/sharing-session/messages-api/mosquitto.conf:/mosquitto/config/mosquitto.conf eclipse-mosquitto
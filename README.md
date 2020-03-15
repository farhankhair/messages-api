# messages-api
Simple API that use Go, MQTT for a simple messaging API

#### How To Run
    1. Ensure Go and MQTT installed on your computer
    2. type `make run` / `go run main.go` to run the program or `make test` / `go test -v ./...` to test the program

#### Simple API

##### 1. API for sending a message Just send one parameter string for message After sending should be get response (REST / GraphQL / etc)
    localhost:8080/message/{message}

##### 2. API for collect message that has been sent out API can get all previously sent messages (REST / GraphQL / etc)
    localhost:8080/message

##### 3. API for display message in real time API should be long live connection to retrieve messages after send at realtime (MQTT / Websocket / etc)
    Server act as a publisher and a subscriber with mqtt server.
    Default address for mqtt    : `mqtt://<username>:<password>@<address>:<port>/<topic>`.
    Address mqtt in mylocal     : `mqtt://mqtt_user_name:mqtt_password@localhost:1883/message-api-realtime`.
    subscriber clientID         : `sub`
    publisher clientID          : `pub`

    When server receiving GET request for sending message, the publisher will publish the body of the message and the subscriber will receive the message.

###### Created By : Farhan Ramadhan Syah Khair
package mqtt

import (
	"fmt"
	"log"
	"net/url"
	"os"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

// Publisher is to create mqtt client as a publisher
func Publisher() mqtt.Client {
	cloudMQTT := os.Getenv("CLOUDMQTT_URL")
	if cloudMQTT == "" {
		cloudMQTT = LocalAddressMQTT()
	}

	uri, err := url.Parse(cloudMQTT)
	if err != nil {
		log.Fatal(err)
	}

	topic := uri.Path[1:len(uri.Path)]
	if topic == "" {
		topic = "message-api-realtime"
	}

	client := Connect("pub", uri)

	return client
}

// Connect :nodoc:
func Connect(clientID string, uri *url.URL) mqtt.Client {
	options := createClientOptions(clientID, uri)
	client := mqtt.NewClient(options)

	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
	return client
}

// Listen is to create subscriber and listening
func Listen(uri *url.URL, topic string) {
	client := Connect("sub", uri)
	client.Subscribe(topic, 0, func(client mqtt.Client, msg mqtt.Message) {
		log.Printf("[%s] : %s\n", msg.Topic(), string(msg.Payload()))
	})
}

func createClientOptions(clientID string, uri *url.URL) *mqtt.ClientOptions {
	options := mqtt.NewClientOptions()
	options.AddBroker(fmt.Sprintf("tcp://%s", uri.Host))
	options.SetUsername(uri.User.Username())
	password, exist := uri.User.Password()
	if !exist {
		log.Fatalf("Password Client Not Exist\n")
	}
	options.SetPassword(password)
	options.SetClientID(clientID)

	return options
}

// LocalAddressMQTT is to get local setup for mqtt
func LocalAddressMQTT() string {
	username := "mqtt_user_name"
	password := "mqtt_password"
	address := "localhost:1883"
	topic := "message-api-realtime"
	cloudMQTT := fmt.Sprintf("mqtt://%s:%s@%s/%s", username, password, address, topic)
	return cloudMQTT
}

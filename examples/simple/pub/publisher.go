package main

import (
	"log"
	"os"

	"github.com/Anth80/nexus/client"
	"github.com/Anth80/nexus/wamp"
)

const (
	addr  = "ws://localhost:8080/ws"
	realm = "realm1"

	exampleTopic = "example.hello"
)

func main() {
	logger := log.New(os.Stdout, "", 0)
	cfg := client.Config{
		Realm:  realm,
		Logger: logger,
	}

	// Connect publisher session.
	publisher, err := client.ConnectNet(addr, cfg)
	if err != nil {
		logger.Fatal(err)
	}
	defer publisher.Close()

	// Publish to topic.
	err = publisher.Publish(exampleTopic, nil, wamp.List{"hello world"}, nil)
	if err != nil {
		logger.Fatal("subscribe error:", err)
	}
	logger.Println("Published", exampleTopic, "event")
}

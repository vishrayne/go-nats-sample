package natscheck

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"runtime"
	"syscall"

	nats "github.com/nats-io/go-nats"
)

// StartSubscriber for given topic
func StartSubscriber(topic string) {
	natsConnection, _ := nats.Connect(nats.DefaultURL)
	log.Println("Connected to ", nats.DefaultURL)

	log.Println("Subscribing to topic, ", topic)
	natsConnection.Subscribe(topic, func(msg *nats.Msg) {
		log.Println("Received message: ", string(msg.Data))
	})

	runtime.Goexit()
}

// StartPublisher for given topic
func StartPublisher(topic string) {
	natsConnection, _ := nats.Connect(nats.DefaultURL)
	defer natsConnection.Close()

	log.Println("Connected to ", nats.DefaultURL)
	log.Println("^C to exit")

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	msg := make(chan string, 1)
	fmt.Print("Message to publish: ")

	go func() {
		var s string
		for {
			fmt.Scan(&s)
			msg <- s
		}
	}()

loop:
	for {
		select {
		case <-sigs:
			log.Println("Exiting...")
			break loop
		case s := <-msg:
			log.Println("publishing ==>", s)
			natsConnection.Publish(topic, []byte(s))
			fmt.Print("Message to publish: ")
		}

	}
}

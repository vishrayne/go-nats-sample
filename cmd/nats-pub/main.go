package main

import (
	natscheck "github.com/vishrayne/go-nats-sample"
)

func main() {
	natscheck.StartPublisher("foo")
}

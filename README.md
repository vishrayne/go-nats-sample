# go-nats-sample
[![baby-gopher](https://raw.githubusercontent.com/drnic/babygopher-site/gh-pages/images/babygopher-logo-small.png)](http://www.babygopher.org)

Attempt to learn how NATS messaging works. There are 2 binaries available under cmd folder
1. nats-pub, a simple NATS publisher.
2. nats-sub, a simple NATS subscriber.

# Prerequisite
You will need the NATS server. Install it using the go tool, \
`$ go get github.com/nats-io/gnatsd`

# Setup

1. Clone the project.
2. `$ dep ensure`
3. To create the required binaries, \
`$ go install github.com/vishrayne/go-nats-sample/cmd/...` 
4. Make sure the NATS server is up by running, \
`$ gnatsd` 

# Usage

1. Start the subscriber in a new terminal 
```
$ nats-sub
  Connected to  nats://localhost:4222
  Subscribing to topic,  foo
```

2. Open another terminal and start the publisher
```
$ nats-pub
  2018/06/05 17:43:04 Connected to  nats://localhost:4222
  2018/06/05 17:43:04 ^C to exit
  Message to publish: hello
  publishing ==> hello
  Message to publish:
  ...
```

Switch back to the previous terminal where you have started the subcsriber to see the incoming messages :)







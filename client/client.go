package main

import (
	"fmt"

	zmq "github.com/pebbe/zmq4"
)

func main() {
	zctx, _ := zmq.NewContext()

	// Socket to talk to server
	fmt.Printf("Connecting to the server...\n")
	s, _ := zctx.NewSocket(zmq.REQ)
	s.Connect("tcp://localhost:8000")

	// Do 10 requests, waiting each time for a response
	for i := 0; i < 5; i++ {
		fmt.Printf("Sending request %d...\n", i)
		s.Send("Request from client", 0)

		message, _ := s.Recv(0)
		fmt.Printf("Received response: %d, %s ]\n", i, message)
	}
}

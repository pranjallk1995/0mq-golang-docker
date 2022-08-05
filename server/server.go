package main

import (
	"log"
	"time"

	zmq "github.com/pebbe/zmq4"
)

func main() {
	zctx, _ := zmq.NewContext()

	s, _ := zctx.NewSocket(zmq.REP)
	s.Bind("tcp://*:8000")

	for {
		// Wait for next request from client
		message, _ := s.Recv(0)
		log.Printf("Received %s\n", message)

		// Do some 'work'
		time.Sleep(time.Second * 1)

		// Send reply back to client
		s.Send("Response from server", 0)
	}
}

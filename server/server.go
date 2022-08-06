package main

import (
	"os"
	"fmt"
	"log"
	"time"

	zmq "github.com/pebbe/zmq4"
)

func main() {
	zctx, _ := zmq.NewContext()

	s, _ := zctx.NewSocket(zmq.REP)
	socket := "tcp://*:" + os.Getenv("PORT")
	s.Bind(socket)
	fmt.Println("Server running on: " + socket)

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

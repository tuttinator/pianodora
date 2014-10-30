package main

import (
	"log"
	"net"
)

type Event struct {
	Name    string `json:"name"`
	Details string `json:"details"`
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	ln, err := net.Listen("tcp", ":9123")
	check(err)

	msgchan := make(chan string)
	go publishMessages(msgchan)

	for {
		conn, err := ln.Accept()
		check(err)

		go handleConnection(conn, msgchan)
	}
}

func handleConnection(c net.Conn, msgchan chan<- string) {
	buf := make([]byte, 4096)

	for {
		n, err := c.Read(buf)
		if err != nil || n == 0 {
			c.Close()
			break
		}
		msgchan <- string(buf[0:n])
	}
	log.Printf("Connection from %v closed.", c.RemoteAddr())

	// Shut down the connection.
	c.Close()
}

func publishMessages(msgchan <-chan string) {
	for msg := range msgchan {
		log.Printf("new message: %s", msg)
	}
}

package main

import (
	"encoding/json"
	"log"
	"net"
	"strings"
)

type PandoraMessage struct {
	Title    string
	Artist   string
	Album    string
	CoverArt string
}

type PianobarEvent struct {
	Name    string
	Details string
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
		s := string(buf[0:n])

		if strings.HasSuffix(s, "\n") {
			go publish(s)
			c.Close()
		}

		msgchan <- s
	}
	log.Printf("Connection from %v closed.", c.RemoteAddr())

	// Shut down the connection.
	c.Close()
}

func publish(s string) PandoraMessage {
	message := parse(s)
	log.Println(message)
	return message
}

func parse(s string) PandoraMessage {
	var event PianobarEvent
	err := json.Unmarshal([]byte(s), &event)
	check(err)

	var message PandoraMessage
	lines := strings.Split(event.Details, "\n")
	for _, line := range lines {
		items := strings.Split(line, "=")
		switch items[0] {
		case "artist":
			message.Artist = items[1]
		case "album":
			message.Album = items[1]
		case "coverArt":
			message.CoverArt = items[1]
		case "title":
			message.Title = items[1]
		}
	}
	return message
}

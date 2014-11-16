package main

import (
	"encoding/json"
	"fmt"
	"net"
	"os"
)

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
	e := &PianobarEvent{Name: extractArgs(os.Args), Details: extractStdin(os.Stdin)}
	payload, err := json.Marshal(e)
	check(err)

	conn, err := net.Dial("tcp", "localhost:9123")
	check(err)

	fmt.Fprintf(conn, "%v\n", string(payload))
}

func extractStdin(file *os.File) string {
	fi, err := file.Stat()
	check(err)

	size := fi.Size()

	if size > 0 {
		// fmt.Printf("%v bytes available in Stdin\n", size)
		data := make([]byte, size)
		file.Read(data)
		check(err)
		return string(data)
	} else {
		return ""
	}
}

func extractArgs(args []string) string {
	if len(args) > 1 {
		return os.Args[1]
	}
	return ""
}

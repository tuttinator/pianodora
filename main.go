package main

import (
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	eventDetails := extractStdin(os.Stdin)
	fmt.Printf(eventDetails)
	eventName := extractArgs(os.Args)
	fmt.Println(eventName)
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

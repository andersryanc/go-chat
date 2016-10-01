package main

import (
	"flag"
	"os"

	"./lib"
)

func main() {
	var isHost bool

	// NOTE - be sure to pass a reference for "isHost" so that it gets updated (and not the copy)
	flag.BoolVar(&isHost, "listen", false, "Listens on the specified ip address")
	flag.Parse()

	if isHost {
		// go run main.go -listen <ip>
		// NOTE - notice that <ip> is the 2nd arg in the run above
		connIP := os.Args[2]
		lib.RunHost(connIP)
	} else {
		// go run main.go <ip>
		// NOTE - notice that <ip> is the 1st arg in the run above
		connIP := os.Args[1]
		lib.RunGuest(connIP)
	}
}

package lib

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

const port = "8080"

// NOTE - public/exported functionss should have a comment

// RunHost takes an ip as an argument and listens for connections on that ip
func RunHost(ip string) {
	ipAndPort := ip + ":" + port

	// create listener
	listener, listenErr := net.Listen("tcp", ipAndPort)
	if listenErr != nil {
		// fmt.Println("Error", listenErr)
		// os.Exit(1)

		// NOTE - this is an easier way to do the above
		log.Fatal("Error: ", listenErr)
	}
	fmt.Println("Listening on: ", ipAndPort)

	// create connection
	conn, acceptErr := listener.Accept()
	if acceptErr != nil {
		log.Fatal("Error: ", acceptErr)
	}
	fmt.Println("New connection accepted")

	// create reader and display messages
	for {
		readMessageFromConn(conn)

		// After receiving message, let the host send one back
		sendMessageToConn(conn)
	}
}

// RunGuest takes a destination ip as an argument and connects to that ip
func RunGuest(ip string) {
	ipAndPort := ip + ":" + port

	// connect to the host
	conn, dialErr := net.Dial("tcp", ipAndPort)
	if dialErr != nil {
		log.Fatal("Error: ", dialErr)
	}

	for {
		sendMessageToConn(conn)

		// After sending a message, wait to see if the host responds
		readMessageFromConn(conn)

	}
}

func readMessageFromConn(conn net.Conn) {
	reader := bufio.NewReader(conn)
	message, replyErr := reader.ReadString('\n')
	if replyErr != nil {
		log.Fatal("Error: ", replyErr)
	}
	fmt.Println("Message recieved:", message)
}

func sendMessageToConn(conn net.Conn) {
	fmt.Print("Send message: ")
	replyReader := bufio.NewReader(os.Stdin)
	replyMessage, replyErr := replyReader.ReadString('\n')
	if replyErr != nil {
		log.Fatal("Error: ", replyErr)
	}

	// Fprint() takes an "io.Writer" interface
	// which includes our connection
	fmt.Fprint(conn, replyMessage)
}

package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

func handleConnection(conn net.Conn) {
	defer conn.Close()

	file, err := os.OpenFile("received_messages.pb", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		log.Printf("Failed to open file: %v", err)
		return
	}
	defer file.Close()

	n, err := io.Copy(file, conn)
	if err != nil {
		log.Printf("Failed to copy data to file: %v", err)
		return
	}

	fmt.Printf("Received and stored %d bytes\n", n)
}

func main () {
	listener, err := net.Listen("tcp", ":9090")
	if err != nil {
		log.Fatalf("Failed to listen on port 9090: %v", err)
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Failed to accept connection: %v", err)
			continue
		}

		go handleConnection(conn)
	}
}
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func handleRequest(conn net.Conn) {
	for {
		message, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			log.Printf("Error occurred while reading in data from [%s]. Reason: %s\n", conn.LocalAddr().String(), err)
			break
		}

		conn.Write([]byte(strings.ToUpper(message)))
		message = strings.Trim(message, "\n")

		if strings.ToLower(message) == "close" || strings.ToLower(message) == "exit" || strings.ToLower(message) == "quit" {
			log.Printf("[%s] closed the connection\n", conn.RemoteAddr().String())
			break
		} else if strings.ToLower(message) == "ping" {
			conn.Write([]byte("PONG\n"))
		} else if message != "" {
			fmt.Printf(">> %s\n", message)
		}
	}
	conn.Close()
}

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal("Failed to start server. Reason:", err)
		return
	}
	defer listener.Close()
	log.Print("Listening on port 8080")

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print("Couldn't accept connection. Reason:", err)
			continue
		}
		log.Printf("[%s] just connected\n", conn.RemoteAddr().String())
		go handleRequest(conn)
	}
}

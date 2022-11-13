package main

import (
	"fmt"
	// Uncomment this block to pass the first stage
	"net"
	"os"
)

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	fmt.Println("Logs from your program will appear here!")

	// Uncomment this block to pass the first stage
	//
	l, err := net.Listen("tcp", "0.0.0.0:6379")
	if err != nil {
		fmt.Println("Failed to bind to port 6379")
		os.Exit(1)
	}
	defer l.Close()

	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Failed to accept connection", err.Error())
			os.Exit(1)
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	for {
		buf := make([]byte, 1024)
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("Failed to read from connection", err.Error())
			return
		}
		fmt.Println("Received", string(buf[:n]))
		// check for the "PING" command in resp protocol
		//if string(buf[:n]) == "*1\r
		//if string(buf[:n]) == "PING" {
		//	fmt.Println("Sending PONG")
		// Send PONG in resp protocol
		conn.Write([]byte("+PONG\r\n"))
		//}
	}
}

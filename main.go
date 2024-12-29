package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func main() {
	PORT := ":6379"
	fmt.Println("Listening on port " + PORT)

	// Create a new server
	l, err := net.Listen("tcp", PORT)
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		return
	}

	conn, err := l.Accept()
	if err != nil {
		fmt.Println("Error accepting: ", err.Error())
		return
	}

	defer conn.Close()

	for {
		buf := make([]byte, 1024)

		// Rread msg from client
		_, err := conn.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println("Error reading:", err.Error())
			os.Exit(1)
		}

		// ignore request and send back a PONG
		conn.Write([]byte("+PONGG\r\n"))
	}
}

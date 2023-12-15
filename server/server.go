package main

import (
	"errors"
	"fmt"
	"io"
	"net"
)

func main() {
	ln, err := net.Listen("tcp", ":8300")
	LogError(err)

	defer ln.Close()

	Pl("Server is listening")

	for {
		conn, err := ln.Accept()
		LogError(err)
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	buffer := make([]byte, 1024)
	for {
		n, err := conn.Read(buffer)
		if err != nil {
			if !errors.Is(err, io.EOF) {
				LogError(err)
			}
			break
		}
		received := string(buffer[:n])
		Pl("Received:", received)

		conn.Write([]byte(fmt.Sprintf("Message Received from %s\n", received)))
	}
	Pl("Ends")
}

func LogError(err error) {
	if err != nil && !errors.Is(err, io.EOF) {
		fmt.Println("\n\n Error thrown is >>> \t\t", err.Error(), "\n\n ")
		panic(err)
	}
}

/** pl: for fmt.Println */
func Pl(msg ...any) {
	fmt.Println("\n", msg, "\n ")
}

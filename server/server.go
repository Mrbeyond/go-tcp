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

	conn, err := ln.Accept()
	LogError(err)
	defer conn.Close()

	Pl("Connected to client!")

	for {
		buffer := make([]byte, 1024)
		n, err := conn.Read(buffer)
		if errors.Is(err, io.EOF) {
			break
		}
		LogError(err)
		Pl("Received: ", string(buffer[:n]))
		conn.Write([]byte("Mssage Received\n"))
	}
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

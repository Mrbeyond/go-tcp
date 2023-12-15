package main

import (
	"fmt"
	"net"
)

func main() {

	conn, err := net.Dial("tcp", "localhost:8300")
	LogError(err)

	defer conn.Close()
	Pl("Connected to server.")

	message := "Hello Server! \tthis is client"
	// tick := time.Tick(time.Second * 1)

	_, err = conn.Write([]byte(message))
	LogError(err)

	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	LogError(err)
	Pl("Response from server >>\t", string(buffer[:n]))

}

func LogError(err error) {
	if err != nil {
		fmt.Println("\n\n Error thrown is >>> \t\t", err.Error(), "\n\n ")
		panic(err)
	}
}

/** pl: for fmt.Println */
func Pl(msg ...any) {
	fmt.Println("\n", msg, "\n ")
}

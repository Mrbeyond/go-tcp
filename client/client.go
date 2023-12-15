package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net"
	"time"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8300")
	LogError(err)
	defer conn.Close()
	Pl("Connected to server.")

	messages := map[string]struct {
		Name    string
		Message string
	}{
		"One":   {Name: "One", Message: "This is from One"},
		"Two":   {Name: "Two", Message: "This is from Two"},
		"Three": {Name: "Three", Message: "This is from Three"},
		"Four":  {Name: "Four", Message: "This is from Four"},
		"Five":  {Name: "Five", Message: "This is from Five"},
	}

	for _, data := range messages {
		payload, err := json.Marshal(data)
		LogError(err)
		_, err = conn.Write(payload)
		Pl(err)

		buffer := make([]byte, 1024)
		n, err := conn.Read(buffer)
		if err != nil && !errors.Is(err, io.EOF) {
			LogError(err)
		}
		Pl("Received response:", string(buffer[:n]))
		time.Sleep(1 * time.Second)
	}
}

func LogError(err error) {
	if err != nil {
		fmt.Println("\n\n Error thrown is >>> \t\t", err.Error(), "\n\n ")
		panic(err)
	}
}

/** pl: for fmt.Println */
func Pl(msg ...interface{}) {
	fmt.Println("\n", msg, "\n ")
}

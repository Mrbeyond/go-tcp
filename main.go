package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"sync"
)

type Client struct {
	conn     net.Conn
	username string
}

var (
	clients     = make(map[string]*Client)
	clientsLock sync.Mutex
)

func broadcastMessage(message string) {
	clientsLock.Lock()
	defer clientsLock.Unlock()

	for _, client := range clients {
		fmt.Fprintf(client.conn, "%s\n", message)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	fmt.Fprintf(conn, "Enter your username: ")
	scanner := bufio.NewScanner(conn)
	scanner.Scan()
	username := scanner.Text()

	newClient := &Client{
		conn:     conn,
		username: username,
	}

	clientsLock.Lock()
	clients[username] = newClient
	clientsLock.Unlock()

	// fmt.Printf("%s joined the chat\n", username)
	broadcastMessage(fmt.Sprintf("%s joined the chat", username))

	for {
		msg, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			break
		}

		msg = strings.TrimSpace(msg)

		if strings.HasPrefix(msg, "@") {
			split := strings.SplitN(msg[1:], " ", 2)
			recipient := split[0]
			content := split[1]

			clientsLock.Lock()
			recipientClient, exists := clients[recipient]
			clientsLock.Unlock()

			if exists {
				fmt.Fprintf(recipientClient.conn, "[DM from %s] %s\n", username, content)
			} else {
				fmt.Fprintf(conn, "User '%s' not found or offline\n", recipient)
			}
		} else {
			broadcastMessage(fmt.Sprintf("[%s] %s", username, msg))
		}
	}

	clientsLock.Lock()
	delete(clients, username)
	clientsLock.Unlock()

	fmt.Printf("%s left the chat\n", username)
	broadcastMessage(fmt.Sprintf("%s left the chat", username))
}

func main() {
	ln, err := net.Listen("tcp", ":8200")
	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}
	defer ln.Close()

	fmt.Println("Server started. Waiting for connections...")

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}

		go handleConnection(conn)
	}
}

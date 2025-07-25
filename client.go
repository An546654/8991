package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Connection error:", err)
		return
	}
	defer conn.Close()

	message, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		fmt.Println("Read error:", err)
		return
	}

	if message == "OK\n" {
		fmt.Println("Success: received", message)
	} else {
		fmt.Println("Unexpected response:", message)
	}
}


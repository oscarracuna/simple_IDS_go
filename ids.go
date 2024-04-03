package main

import (
	"fmt"
	"net"

)

func main() {
	listenAddress := "0.0.0.0:8080"

	listener, err := net.Listen("tcp", listenAddress)
	if err != nil {
		fmt.Println("Error listening: ", err.Error())
		return

	}

	defer listener.Close()

	fmt.Println("Listening incoming connections... using port ", listenAddress)
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection...", err.Error())
			continue

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	buffer := make([]byte, 1024)
	if err != nil {
		fmt.Println("Error reading the data...", err.Error())
		return
	}

	processData(buffer[:bytesRead])
}

func processData(data []byte) {
	//Here goes the ids logic

	//This is for testing purposes only
	//It'll just print the data received

	fmt.Println("Data received... ", string(data))
	//Here goes the rules and actions to take


}

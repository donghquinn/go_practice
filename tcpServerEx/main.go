package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	server, error := net.Listen("tcp", ":8080")

	if error != nil {
		fmt.Println(error)
		log.Fatal(error)
	}

	for {
		connect, error := server.Accept()

		if error != nil {
			fmt.Println(error)
			log.Fatal(error)
		}

		go func(connection net.Conn) {
			io.Copy(connection, connection)

			connection.Close()
		}(connect)
	}
}

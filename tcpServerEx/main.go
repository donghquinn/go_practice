package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {
	fmt.Println("Server Listening Start")

	// server 시작
	server, error := net.Listen("tcp", ":8080")

	// 에러 핸들링
	if error != nil {
		fmt.Println(error)
		log.Fatal(error)
	}

	// 취소 지연
	defer server.Close()

	for {
		// 서버가 accept 한 동안
		connect, error := server.Accept()

		fmt.Println("Server is Listening : 8080")

		if error != nil {
			fmt.Println(error)
			log.Fatal(error)
		}

		// handle 메서드 호출
		go handle(connect)
	}
}

func handle(connect net.Conn) {
	var n int

	scan := bufio.NewScanner(connect)

	for scan.Scan() {
		// 스캔에 호출된 가장 최근에 생성된 토큰 리턴
		line := scan.Text()

		if n == 0 {
			fmt.Println(line)

			stringFields := strings.Fields(line)

			method := stringFields[0]
			route := stringFields[1]

			switch {
			case method == "GET" && route == "/":
				println(method, route)

			case method == "GET" && route == "/dog":
				println(method, route)
			}

		}

		if line == "" {
			fmt.Println("scanned text is empty")
			break
		}

		n++

	}
	fmt.Println("Server Stopping")
	connect.Close()
}

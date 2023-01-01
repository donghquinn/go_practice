package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Server is Starting")

	// 라우트 / 에 대한 요청이 들어왔을 때, handler 함수 호출
	http.HandleFunc("/", handler)

	// 서버는 8080 포트 - ListenAndServe는 가장 마지막 줄에
	http.ListenAndServe(":8080", nil)
}

func handler(response http.ResponseWriter, request *http.Request) {
	fmt.Println("Got Request To / route")
	// 응답 값 인코딩하여 응답.
	json.NewEncoder(response).Encode(`Received Request`)
}

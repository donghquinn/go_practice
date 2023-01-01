package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

/*
참고: https://github.com/SimonWaldherr/golang-examples/blob/master/advanced/httpupload.go
*/

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

// 파일 업로드 요청 라우트
func fileUpHandler(response http.ResponseWriter, request *http.Request) {
	fmt.Println("Got Request to upload route")

	// 요청 방식이 POST가 아닐 때 요청 거절 - 얼리 리턴
	if request.Method != "POST" {
		log.Fatal("Request Methos Is Not POST")

		// 응답 인코딩
		json.NewEncoder(response).Encode("Request Methos Is Not POST")

		return
	}

	reader, err := request.MultipartReader()

	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)

		return
	}
}

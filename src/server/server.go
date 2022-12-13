package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	_ "strings"

	"log"
	// "strings"

	mux "github.com/gorilla/mux"
)

// 요청으로 들어올 유저 정보

const port = ":3000"

func main() {
	router := mux.NewRouter()
	fmt.Println("Listening on: ", port)

	router.HandleFunc("/admin/count", handleCount)
}

func postRequest(request *http.Request) {
	request.Header.Set("Content-Type", "x-www-form-urlencoded")

	bodyData := request.Body
}

func handleCount(response http.ResponseWriter, request *http.Request) {

	// fmt.Println("Received Request: ", request.Body.Read("", response)
	log.Default()

	json.NewEncoder(response).Encode(`Received Request: {request}`)
}

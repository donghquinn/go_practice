package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/result", resultHandler)

	fmt.Println("Server Start")
	// server start
	http.ListenAndServe(":8080", nil)

}

func resultHandler(response http.ResponseWriter, request *http.Request) {
	json.NewEncoder(response).Encode("Hi There")

	fmt.Println("Responded")

}

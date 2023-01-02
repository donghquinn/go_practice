package main

import (
	"encoding/json"
	"net/http"
)

func main() {
	http.HandleFunc("/result", resultHandler)

	// server start
	http.ListenAndServe(":8080", nil)

}

func resultHandler(response http.ResponseWriter, request *http.Request) {
	json.NewEncoder(response).Encode("Hi There")
}

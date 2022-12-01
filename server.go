package main

import (
	"net/http"
)

// 요청으로 들어올 유저 정보
type User struct {
	name   string
	age    string
	email  string
	birth  string
	gender string
}

func postUser() {
	const exampleUrl = "https://www.example.com"
	const userInfo = http.Post(exampleUrl)
}

func apiListening() {
	const port = ":8080"

	http.ListenAndServe(port, nil)
}
